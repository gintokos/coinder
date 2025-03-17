package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	appcfg "github.com/gintokos/coinder/backend/internal/config"
	"github.com/gintokos/coinder/backend/internal/middleware"
	"github.com/gintokos/coinder/backend/internal/storage"
	"github.com/gintokos/coinder/backend/pkg/sl"
	"github.com/gintokos/coinder/backend/pkg/telegram"
	"github.com/spf13/viper"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

const (
	API_ROUTE = "/api/v1"
)

type App struct {
	bot *telegram.Bot

	context   context.Context
	ctxcancel context.CancelFunc
	domain    string

	*userServiceProvider
	*coinServiceProvider
	*telegraWebhookServiceProvider

	server   *http.Server
	api      *gin.RouterGroup
	router   *gin.Engine
	database *storage.Storage
}

func NewApp(db *storage.Storage) (*App, error) {
	a := App{
		database: db,
	}
	a.context, a.ctxcancel = context.WithCancel(context.Background())

	return &a, a.initDeps()
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initDomain,

		a.initRouter,

		a.initServer,

		a.initBot,

		a.initUserServiceProvider,

		a.initCoinServiceProvider,

		a.initTelegramWebhookServiceProvider,
	}

	for _, f := range inits {
		err := f()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initDomain() error {
	if viper.GetString("env") == appcfg.LOCAL_WITH_NGROK {
		a.domain = viper.GetString("ngrok_domain")
	} else {
		a.domain = viper.GetString("domain")
	}

	return nil
}

func (a *App) initBot() error {
	slog.Debug("initBot with domain: " + a.domain)
	bot, err := telegram.NewBot(a.context,
		viper.GetString("botToken"),
		a.domain,
		"coinder",
		API_ROUTE+"/telegramwebhook",
	)
	a.bot = bot
	if err != nil {
		slog.Error("creating telegram bot was failed", sl.Err(err))
		return err
	}
	slog.Info("Telegram bot was created")
	return nil
}

func (a *App) initRouter() error {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// in local with ngrok redirect all requests to http://localhost:5173 where should be running vite dev server
	env := viper.GetString("env")

	if env == appcfg.LOCAL || env == appcfg.LOCAL_WITH_NGROK {
		r.Use(func(c *gin.Context) {
			if strings.HasSuffix(c.Request.URL.Path, ".js") {
				c.Header("Content-Type", "application/javascript; charset=utf-8")
				if strings.Contains(c.Request.URL.Path, "index-") || strings.Contains(c.Request.URL.Path, "chunk-") {
					c.Writer.Header().Set("Content-Type", "text/javascript; charset=utf-8")
				}
			}
			c.Next()
		})
		r.NoRoute(func(c *gin.Context) {
			proxy := &httputil.ReverseProxy{
				Director: func(req *http.Request) {
					req.URL.Scheme = "http"
					req.URL.Host = "localhost:5173"
				},
			}
			proxy.ServeHTTP(c.Writer, c.Request)
		})
	}

	a.api = r.Group(API_ROUTE)
	a.api.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})

	// logging middleware
	a.api.Use(middleware.LoggingMiddleware())

	// auth middleware and handler
	token := viper.GetString("botToken")
	cookieName := "ta_t"
	if env != appcfg.LOCAL {
		a.api.POST("/auth", telegram.AuthHandler(token, time.Hour*24*7, 24, cookieName, a.domain, false))
		a.api.Use(telegram.AuthMiddleware(cookieName, token))
	} else {
		// local sets standart user for all reqs
		a.api.Use(telegram.TestMiddleware(telegram.TauthData{
			ID:        128389,
			FirstName: "test_user",
			Username:  "test_username",
			PhotoURL:  "https://t.me/i/userpic/320/GfOke9XNyVTOgIQ7JMewCAqQ6oqOfLNzDRAWf4DxhyQ.jpg",
			AuthDate:  1238615,
			Hash:      "randomhash",
		}))
	}
	// refresh token
	a.api.GET("/auth/refresh", telegram.RefreshTokenHandler(token, time.Hour*24*7, cookieName, a.domain, false))
	// endpoint checks is authed
	a.api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pong": "success",
		})
	})

	a.router = r
	slog.Info("Router was created")
	return nil
}

func (a *App) initUserServiceProvider() error {
	a.userServiceProvider = newUserServiceProvider(a.api, a.bot, a.database)
	slog.Info("UserServiceProvider was created")
	return nil
}

func (a *App) initCoinServiceProvider() error {
	a.coinServiceProvider = newCoinServiceProvider(a.api, a.database)
	slog.Info("CoinServiceProvider was created")
	return nil
}

func (a *App) initTelegramWebhookServiceProvider() error {
	a.telegraWebhookServiceProvider = newTelegramWebhookServiceProvider(a.router, a.bot, a.database)
	slog.Info("TelegramWebhookServiceProvider was created")
	return nil
}

func (a *App) initServer() error {
	server := &http.Server{
		Handler:           a.router,
		Addr:              "0.0.0.0:8080",
		ReadTimeout:       viper.GetDuration("server.readTimeout"),
		WriteTimeout:      viper.GetDuration("server.writeTimeout"),
		IdleTimeout:       viper.GetDuration("server.idleTimeout"),
		ReadHeaderTimeout: viper.GetDuration("server.readHeaderTimeout"),
	}
	slog.Debug("Server inited", slog.Attr{
		Key:   "addr",
		Value: slog.StringValue(server.Addr),
	})
	a.server = server

	slog.Info("Server was created")

	return nil
}

func (a *App) MustRun() error {
	env := viper.GetString("env")

	switch env {
	case appcfg.LOCAL_WITH_NGROK:
		tun, err := ngrok.Listen(a.context, config.HTTPEndpoint(
			config.WithDomain(viper.GetString("ngrok_domain")),
		), ngrok.WithAuthtoken(viper.GetString("ngrok_token")))
		if err != nil {
			slog.Error("Listening ngrok server was failed", sl.Err(err))
			os.Exit(1)
		}
		slog.Debug("ngrok tunnel was created")

		a.startServer(tun)

	default:
		a.startServer()
	}

	a.startBot()
	slog.Info("App started work")

	return nil
}

func (a *App) startServer(listeners ...net.Listener) {
	if len(listeners) > 0 {
		tun := listeners[0]
		go func() {
			if err := a.server.Serve(tun); err != nil && !errors.Is(err, http.ErrServerClosed) {
				slog.Error("Starting server with ngrok was failed", sl.Err(err))
				os.Exit(1)
			}
		}()
		slog.Debug("server was started with ngrok")
		localServer := http.Server{
			Handler: a.router,
			Addr:    "127.0.0.1:8081",
		}
		go func() {
			err := localServer.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				slog.Error("error on hosting local server 127.0.0.1:8081", sl.Err(err))
				os.Exit(1)
			}
		}()
		slog.Debug("local server was created")
		return
	}

	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Starting server local was failed", sl.Err(err))
			os.Exit(1)
		}
	}()

}

func (a *App) startBot() {
	go a.bot.MustStart()
}

func (a *App) GraceFullShutDown(ctx context.Context) error {
	var servererr, dberr error
	if err := a.server.Shutdown(ctx); err != nil {
		servererr = err
	}

	err := a.database.GraceFullShutDown()
	if err != nil {
		dberr = err
	}

	if servererr != nil || dberr != nil {
		return fmt.Errorf("error on shutting down server: %v, error on shutting down db: %v", servererr, dberr)
	}

	a.ctxcancel()
	return nil
}

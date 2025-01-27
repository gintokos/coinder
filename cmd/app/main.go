package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gintokos/coinder/internal/app"
	"github.com/gintokos/coinder/internal/config"
	"github.com/gintokos/coinder/internal/storage/postgres"
	"github.com/gintokos/coinder/pkg/sl"
)

func main() {
	config.MustInitForServer()

	db, err := postgres.NewDatabase()
	if err != nil {
		slog.Error("error on creating database", sl.Err(err))
		os.Exit(1)
	}

	app, err := app.NewApp(db)
	if err != nil {
		slog.Error("error on creating app", sl.Err(err))
		os.Exit(1)
	}

	go func() {
		if err := app.Run(); err != nil {
			slog.Error("error on running app", sl.Err(err))
			os.Exit(1)
		}
	}()

	closed := make(chan os.Signal, 1)

	signal.Notify(closed, os.Interrupt, syscall.SIGTERM)

	<-closed

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.GraceFullShutDown(ctx); err != nil {
		slog.Error("error on gracefull shut down", sl.Err(err))
	} else {
		slog.Info("app shut down successfully")
	}
}

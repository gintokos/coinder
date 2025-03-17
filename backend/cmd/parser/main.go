package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gintokos/coinder/backend/internal/config"
	"github.com/gintokos/coinder/backend/internal/parser"
	"github.com/gintokos/coinder/backend/internal/storage/postgres"
	"github.com/gintokos/coinder/backend/pkg/sl"
	"github.com/spf13/viper"
)

func main() {
	config.MustInitForParser()

	// wait for 10 seconds to allow the database to start from backend
	time.Sleep(10 * time.Second)

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	db, err := postgres.NewDatabase()
	if err != nil {
		slog.Error("error on creating database: ", sl.Err(err))
		os.Exit(1)
	}
	slog.Info("Database was created")

	parsCfg := parser.Config{
		Timestamp:           viper.GetDuration("parser.timestamp"),
		TimeoutForReq:       viper.GetDuration("parser.timeout_for_req"),
		ApiKeyCoinMarketCap: viper.GetString("parser.cmc_apikey"),
	}
	pars := parser.New(parsCfg, db)
	slog.Info("Parser was created")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	go func() {
		err = pars.Run(ctx, &wg)
		if err != nil {
			slog.Error("failed to start parser ", sl.Err(err))
			os.Exit(1)
		}
	}()
	slog.Info("Parser start running")

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, os.Interrupt)

	<-s
	cancel()

	wg.Wait()
	slog.Info("Parser ended work")
}

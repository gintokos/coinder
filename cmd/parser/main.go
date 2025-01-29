package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gintokos/coinder/internal/config"
	"github.com/gintokos/coinder/internal/parser"
	"github.com/gintokos/coinder/internal/storage/postgres"
	"github.com/gintokos/coinder/pkg/sl"
)

func main() {
	config.MustInitForParser()

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
		Timestamp:           time.Second * 10,
		TimeoutForReq:       time.Second * 10,
		ApiKeyCoinMarketCap: "adb5310b-ece6-40c1-9904-caa8f3cc704e",
	}
	pars := parser.NewParser(parsCfg, db)
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

package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gintokos/coinder/internal/config"
	"github.com/gintokos/coinder/internal/parser"
	"github.com/gintokos/coinder/internal/storage/postgres"
	"github.com/gintokos/coinder/pkg/sl"
)

func main() {
	config.MustInitForParser()

	db, err := postgres.NewDatabase()
	if err != nil {
		slog.Error("error on creating database", sl.Err(err))
		os.Exit(1)
	}

	pars, err := parser.NewParser(db)
	if err != nil {
		slog.Error("error on creating parser", sl.Err(err))
		os.Exit(1)
	}

	go func() {
		err := pars.Run()
		if err != nil {
			slog.Error("error on running parser", sl.Err(err))
			os.Exit(1)
		}
	}()

	closed := make(chan os.Signal, 1)

	signal.Notify(closed, os.Interrupt, syscall.SIGTERM)

	<-closed

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pars.GraceFullShutDown(ctx); err != nil {
		slog.Error("error on shutting down parser", sl.Err(err))
		os.Exit(1)
	} else {
		slog.Info("parser stopped successfully")
	}

}

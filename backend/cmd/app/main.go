package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gintokos/coinder/backend/app"
	"github.com/gintokos/coinder/backend/config"
	"github.com/gintokos/coinder/backend/storage"
	"github.com/gintokos/coinder/backend/pkg/sl"
)

func main() {
	config.MustInitForApp()

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Debug("default slog was updated")

	db, err := storage.NewStorage()
	if err != nil {
		slog.Error("error on creating database", sl.Err(err))
		os.Exit(1)
	}
	slog.Info("database was created")

	app, err := app.NewApp(db)
	if err != nil {
		slog.Error("error on creating app", sl.Err(err))
		os.Exit(1)
	}
	slog.Info("app was created")

	app.MustRun()
	slog.Info("app running")

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

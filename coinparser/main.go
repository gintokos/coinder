package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gintokos/coinder/backend/pkg/sl"
	"github.com/gintokos/coinder/coinparser/parser"
	"github.com/gintokos/coinder/coinparser/server"
	"github.com/gintokos/coinder/coinparser/storage"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Debug("default slog was updated")

	// init database for microservice
	db, err := storage.NewStorage()
	if err != nil {
		slog.Error("error on creating database", sl.Err(err))
		os.Exit(1)
	}
	slog.Info("database was created")	
	// end of database init

	// init server for update on direct reqs to coin to keep actual data for coins
	server, err := server.NewServer(db)
	if err != nil {
		slog.Error("error on creating server", sl.Err(err))
		os.Exit(1)
	}
	slog.Info("server was created")

	go func() {
		go server.MustRun()
	}()
	slog.Info("server running")
	// end of server init

	// init parser for keep actual data for all coins every day and get new ones
	parser := parser.New(db)
	slog.Info("parser was created")

	parserctx, cancelparser := context.WithCancel(context.Background())
	go func() {
		go parser.Run(parserctx)
	}()
	slog.Info("parser running")
	// end of parser init

	// Preventing application from shutting down instantly
	closed := make(chan os.Signal, 1)
	signal.Notify(closed, os.Interrupt, syscall.SIGTERM)
	<-closed
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// stoping parser
	cancelparser()

	// stoping server
	if err := server.GraceFullShutDown(ctx); err != nil {
		slog.Error("error on gracefull shut down", sl.Err(err))
	} else {
		slog.Info("shut down successfully")
	}
}

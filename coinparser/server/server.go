package server

import (
	"context"
	"log/slog"

	"github.com/gintokos/coinder/backend/pkg/sl"
	"github.com/gintokos/coinder/coinparser/storage"
)

type Server struct {
	db *storage.Storage
}

func NewServer(storage *storage.Storage) (*Server, error) {
	return &Server{
		db: storage,
	}, nil
}

func (s *Server) MustRun() {
	err := s.Run()
	if err != nil {
		slog.Error("error on running server", sl.Err(err))
		panic(err)
	}
}

func (s *Server) Run() error {
	return nil
}

func (s *Server) GraceFullShutDown(ctx context.Context) error {
	return nil
}

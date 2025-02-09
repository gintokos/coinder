package storage

import (
	"github.com/gintokos/coinder/internal/storage/postgres"
)

type Storage struct {
	*postgres.Database
}

func NewStorage() (*Storage, error) {
	db, err := postgres.NewDatabase()
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}


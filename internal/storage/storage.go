package storage

import (
	"github.com/gintokos/coinder/internal/models"
	"github.com/gintokos/coinder/internal/storage/cache"
	"github.com/gintokos/coinder/internal/storage/postgres"
)

type Storage struct {
	cache *cache.Cache
	*postgres.Database
}

func NewStorage() (*Storage, error) {
	db, err := postgres.NewDatabase()
	if err != nil {
		return nil, err
	}

	c, err := cache.New(db)
	if err != nil {
		return nil, err
	}

	return &Storage{c, db}, nil
}

func (s *Storage) GraceFullShutDown() error {
	return s.cache.GraceFullShutDown()
}

func (s *Storage) DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.DBCoin, error) {
	if opt.LikedToday {
	}

	return s.Database.DefaultSearchCoins(opt)
}

func (s *Storage) ChangeLike(isIncrement bool, coinid int, userid int64) error {
	s.cache.ChangeLike(isIncrement, coinid, userid)
	return nil
}

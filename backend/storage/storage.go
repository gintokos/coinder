package storage

import (
	"github.com/gintokos/coinder/backend/models"
	"github.com/gintokos/coinder/backend/storage/cache"
	"github.com/gintokos/coinder/backend/storage/postgres"
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

	c, err := cache.New()
	if err != nil {
		return nil, err
	}

	return &Storage{c, db}, nil
}

func (s *Storage) GraceFullShutDown() error {
	return s.cache.GraceFullShutDown()
}

func (s *Storage) DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.CoinResp, error) {
	var rcoins []models.CoinResp
	coins, err := s.Database.DefaultSearchCoins(opt)
	if err != nil {
		return nil, err
	}

	if opt.LikedByUser {
		rcoins = s.cache.LikesInfo(coins, opt.UserIDLClient)
		return rcoins, nil
	}

	for _, c := range coins {
		rcoins = append(rcoins, models.CoinResp{
			DBCoin:  c,
			IsLiked: false,
		})
	}

	return rcoins, nil
}

func (s *Storage) ChangeLike(isIncrement bool, coinid int, userid int64) error {
	s.cache.StoreLiked(isIncrement, coinid, userid)
	if isIncrement {
		return s.IncrementLike(coinid, userid)
	} else {
		return s.DecrementLike(coinid, userid)
	}
}

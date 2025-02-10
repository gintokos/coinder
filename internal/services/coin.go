package services

import (
	"errors"
	"log/slog"
	"strconv"
	"time"

	"github.com/gintokos/coinder/internal/handlers"
	"github.com/gintokos/coinder/internal/models"
	"github.com/gintokos/coinder/internal/parser"
	"github.com/gintokos/coinder/pkg/sl"
	"gorm.io/gorm"
)

var _ handlers.CoinService = (*CoinService)(nil)

const (
	PARSER_TIMEOUT = 1 * time.Second
)

type CoinStorage interface {
	UpdateCoins(coin []models.DBCoin) error
	DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.DBCoin, error)
	CustomSearchCoins(opt models.SearchCoinOpt) ([]models.DBCoin, error)
}

type CoinService struct {
	parser.Parser
	storage CoinStorage
}

func NewCoinService(storage CoinStorage) *CoinService {
	ps := parser.NewDefault()

	return &CoinService{
		Parser:  ps,
		storage: storage,
	}
}

func (s *CoinService) DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.DBCoin, error) {
	dbcoins, err := s.storage.DefaultSearchCoins(opt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		slog.Error("error in getting coins from database: ", sl.Err(err))
		return nil, ErrServer
	}

	s.Update(dbcoins)

	return dbcoins, nil
}

func (s *CoinService) Update(dbcoins []models.DBCoin) {
	timer := time.NewTimer(PARSER_TIMEOUT)
	done := make(chan struct{})
	success := false

	ids := make([]string, 0, len(dbcoins))
	for _, coin := range dbcoins {
		ids = append(ids, strconv.Itoa(coin.ID))
	}

	go func() {
		defer close(done)

		parscoins, err := s.Parser.GetListWithoutMeta(ids)
		if err != nil || len(parscoins) == 0 {
			slog.Error("error on getting coins from parser: ", sl.Err(err))
			return
		}

		for i, pscoin := range parscoins {
			dbcoins[i] = models.ToDBcoin(&pscoin)
		}
		success = true
	}()

	select {
	case <-timer.C:
		slog.Warn("parser timeout exceeded")
		return
	case <-done:
		timer.Stop()
	}

	if success {
		go func() {
			err := s.storage.UpdateCoins(dbcoins)
			if err != nil {
				slog.Error("error on updating coins: ", sl.Err(err))
			}
		}()
	}
}

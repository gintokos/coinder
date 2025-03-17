package services

import (
	"log/slog"

	"github.com/gintokos/coinder/backend/handlers"
	"github.com/gintokos/coinder/backend/models"
	"github.com/gintokos/coinder/backend/pkg/sl"
)

var _ handlers.CoinService = (*CoinService)(nil)

type CoinStorage interface {
	UpdateCoins(coin []models.DBCoin) error
	DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.CoinResp, error)
	ChangeLike(isIncrement bool, coinid int, userid int64) error
}

type CoinService struct {
	// parser.Parser
	storage CoinStorage
}

func NewCoinService(storage CoinStorage) *CoinService {
	// ps := parser.NewDefault(viper.GetString("parser.cmc_apikey"))

	return &CoinService{
		// Parser:  ps,
		storage: storage,
	}
}

func (s *CoinService) DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.CoinResp, error) {
	dbcoins, err := s.storage.DefaultSearchCoins(opt)
	if err != nil {
		slog.Error("error in getting coins from database: ", sl.Err(err))
		return nil, err
	}

	// s.Update(dbcoins)

	return dbcoins, nil
}

func (s *CoinService) ChangeLike(isIncrement bool, coinid int, userid int64) error {
	return s.storage.ChangeLike(isIncrement, coinid, userid)
}


// func (s *CoinService) Update(dbcoins []models.DBCoin) {
// 	if len(dbcoins) == 0 {
// 		return
// 	}

// 	timer := time.NewTimer(PARSER_TIMEOUT)
// 	done := make(chan struct{})
// 	success := false

// 	ids := make([]string, 0, len(dbcoins))
// 	for _, coin := range dbcoins {
// 		ids = append(ids, strconv.Itoa(coin.ID))
// 	}

// 	go func() {
// 		defer close(done)

// 		parscoins, err := s.Parser.GetListWithoutMeta(ids)
// 		if err != nil || len(parscoins) == 0 {
// 			slog.Error("error on getting coins from parser for req: ", sl.Err(err))
// 			return
// 		}

// 		for i, pscoin := range parscoins {
// 			dbcoins[i] = models.ToDBcoin(&pscoin)
// 		}
// 		success = true
// 	}()

// 	select {
// 	case <-timer.C:
// 		slog.Warn("parser timeout exceeded")
// 		return
// 	case <-done:
// 		timer.Stop()
// 	}

// 	if success {
// 		go func() {
// 			err := s.storage.UpdateCoins(dbcoins)
// 			if err != nil {
// 				slog.Error("error on updating coins: ", "error", gerror.FullError(err))
// 			}
// 		}()
// 	}
// }


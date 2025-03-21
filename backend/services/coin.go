package services

import (
	"context"
	"log/slog"
	"time"

	"github.com/gintokos/coinder/backend/handlers"
	"github.com/gintokos/coinder/backend/models"
	"github.com/gintokos/coinder/backend/pkg/sl"
	pb "github.com/gintokos/coinder/protos/coinupdateprotos"
)

var _ handlers.CoinService = (*CoinService)(nil)

type CoinStorage interface {
	DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.CoinResp, error)
	ChangeLike(isIncrement bool, coinid int, userid int64) error
}

type CoinService struct {
	pbclient pb.CoinServiceClient
	storage  CoinStorage
}

func NewCoinService(storage CoinStorage, pbclient pb.CoinServiceClient) *CoinService {
	return &CoinService{
		pbclient: pbclient,
		storage:  storage,
	}
}

func (s *CoinService) DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.CoinResp, error) {
	rcoins, err := s.storage.DefaultSearchCoins(opt)
	if err != nil {
		slog.Error("error in getting coins from database: ", sl.Err(err))
		return nil, err
	}

	dbcoins := make([]models.DBCoin, len(rcoins))
	likedStatus := make(map[int]bool, len(rcoins))

	for i, rcoin := range rcoins {
		dbcoins[i] = rcoin.DBCoin
		likedStatus[rcoin.ID] = rcoin.IsLiked
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	pbcoins, err := s.pbclient.UpdateCoins(ctx, models.ToPBCoin(dbcoins))
	if err != nil {
		slog.Error("Error on updating coins from service", sl.Err(err))
		return rcoins, nil
	}

	updatedDBcoins := models.ToDBCoinsFromPB(pbcoins)

	updatedCoinsMap := make(map[int]models.DBCoin, len(updatedDBcoins))
	for _, coin := range updatedDBcoins {
		updatedCoinsMap[coin.ID] = coin
	}

	for i := range rcoins {
		if updatedCoin, exists := updatedCoinsMap[rcoins[i].ID]; exists {
			rcoins[i].DBCoin = updatedCoin
		}
	}

	return rcoins, nil
}

func (s *CoinService) ChangeLike(isIncrement bool, coinid int, userid int64) error {
	return s.storage.ChangeLike(isIncrement, coinid, userid)
}

package services

import (
	"github.com/gintokos/coinder/internal/models"
)

type CoinStorage interface {
	UpdateCoins(coin []models.DBCoin) error
	Coins() ([]models.DBCoin, error)
}

type CoinService struct {
	storage CoinStorage
}

func NewCoinService(storage CoinStorage) CoinService {
	return CoinService{storage: storage}
}

func (s CoinService) Coins(opt models.SearchCoinOpt) ([]models.DBCoin, error) {

}

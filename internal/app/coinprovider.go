package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/handlers"
	"github.com/gintokos/coinder/internal/services"
	"github.com/gintokos/coinder/internal/storage"
)

type coinServiceProvider struct {
	handlers.CoinHandler
	*services.CoinService
	*storage.Storage
}

func newCoinServiceProvider(router *gin.RouterGroup, st *storage.Storage) *coinServiceProvider {
	service := services.NewCoinService(st)
	handler := handlers.NewCoinHandler(router, service)
	return &coinServiceProvider{
		CoinHandler: handler,
		CoinService: service,
		Storage:     st,
	}
}

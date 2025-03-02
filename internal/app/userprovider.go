package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/handlers"
	"github.com/gintokos/coinder/internal/services"
	"github.com/gintokos/coinder/internal/storage"
	"github.com/gintokos/coinder/pkg/telegram"
)

type userServiceProvider struct {
	handlers.UserHandler
	*services.UserService
	*storage.Storage
}

func newUserServiceProvider(router *gin.RouterGroup, bot *telegram.Bot, st *storage.Storage) *userServiceProvider {
	service := services.NewUserService(bot, st)
	handler := handlers.NewUserHandler(router, service)
	return &userServiceProvider{
		UserHandler: handler,
		UserService: service,
		Storage:     st,
	}
}

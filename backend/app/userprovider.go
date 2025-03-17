package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/backend/handlers"
	"github.com/gintokos/coinder/backend/services"
	"github.com/gintokos/coinder/backend/storage"
	"github.com/gintokos/coinder/backend/pkg/telegram"
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

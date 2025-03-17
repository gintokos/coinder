package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/backend/internal/handlers"
	"github.com/gintokos/coinder/backend/internal/services"
	"github.com/gintokos/coinder/backend/internal/storage"
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

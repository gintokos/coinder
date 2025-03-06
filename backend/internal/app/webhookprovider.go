package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/handlers"
	"github.com/gintokos/coinder/internal/services"
	"github.com/gintokos/coinder/internal/storage"
	"github.com/gintokos/coinder/pkg/telegram"
)

type telegraWebhookServiceProvider struct {
	handlers.TelegramWebhookHandler
	*services.TelegramWebhookService
	*storage.Storage
}

func newTelegramWebhookServiceProvider(router *gin.Engine, bot *telegram.Bot, st *storage.Storage) *telegraWebhookServiceProvider {
	service := services.NewTelegramWebhookService(bot, st)
	handler := handlers.NewTelegramWebhookHandler(router, service)
	return &telegraWebhookServiceProvider{
		TelegramWebhookHandler: handler,
		TelegramWebhookService: service,
		Storage:                st,
	}
}

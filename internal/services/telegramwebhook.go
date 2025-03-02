package services

import "github.com/gintokos/coinder/pkg/telegram"

type TelegramWebhookService struct {
	bot     *telegram.Bot
	storage TelegramWebHookStorage
}

type TelegramWebHookStorage interface {
	// CacheProcessingDonate(userid int64, amount int)
	// SaveDonate(userid int64, amount int)
}

func NewTelegramWebhookService(bot *telegram.Bot, storage TelegramWebHookStorage) *TelegramWebhookService {
	return &TelegramWebhookService{
		bot:     bot,
		storage: storage,
	}
}


package handlers

import "github.com/gin-gonic/gin"

type TelegramWebhookHandler struct {
	service TelegramWebhookService
}

type TelegramWebhookService interface {
	
}

func NewTelegramWebhookHandler(router *gin.Engine, service TelegramWebhookService) TelegramWebhookHandler {
	h := TelegramWebhookHandler{
		service: service,
	}
	h.RegisterRoutes(router)
	return h
}

func (h *TelegramWebhookHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/telegramwebhook", h.HandleReq)
}

func (h *TelegramWebhookHandler) HandleReq(c *gin.Context) {
	
}

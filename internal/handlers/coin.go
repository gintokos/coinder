package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/models"
)

type CoinService interface {
	Coins(opt models.SearchCoinOpt) ([]models.DBCoin, error)
}

type CoinHandler struct {
	service CoinService
}

func NewCoinHandler(router *gin.RouterGroup, service CoinService) CoinHandler {
	h := CoinHandler{service}
	h.RegisterRoutes(router)
	return h
}

func (h *CoinHandler) RegisterRoutes(router *gin.RouterGroup) {
	coingrp := router.Group("/coins")

	coingrp.POST("/", h.Coins)
}

func (h *CoinHandler) Coins(c *gin.Context) {
	slog.Info("coin.coins")

	user := UserFromClaims(c)

	var opt models.SearchCoinOpt
	err := c.BindJSON(opt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid options data",
		})
		return
	}

	opt.UserID = user.ID
	coins, err := h.service.Coins(opt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   coins,
	})
}

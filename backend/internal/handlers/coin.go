package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/backend/internal/middleware"
	"github.com/gintokos/coinder/backend/internal/models"
	"github.com/go-playground/validator/v10"
)

type likeReq struct {
	CoindID int `json:"coin_id"`
}

type CoinService interface {
	DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.CoinResp, error)
	ChangeLike(isIncrement bool, coinid int, userid int64) error
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

	coingrp.POST("/default", h.Coins)
	coingrp.POST("/like", h.changeLike(true))
	coingrp.POST("/dislike", h.changeLike(false))
}

func (h *CoinHandler) changeLike(isIncrement bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Info("handlers.coin.changeLike")
		user := middleware.UserFromClaims(c)

		var req likeReq
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "invalid id format",
			})
			return
		}
		err = h.service.ChangeLike(isIncrement, req.CoindID, user.ID)

		if err != nil {
			ErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"status": "ok",
		})
	}
}

func (h *CoinHandler) Coins(c *gin.Context) {
	slog.Info("handlers.coin.Default")

	var sOptReq models.SearchCoinOptReq

	err := c.BindJSON(&sOptReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "invalid search options",
		})
		return
	}

	validate := validator.New()
	if err = validate.Struct(sOptReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	user := middleware.UserFromClaims(c)
	sOpt := models.SearchCoinOpt{
		UserIDLClient:    user.ID,
		SearchCoinOptReq: sOptReq,
	}

	coins, err := h.service.DefaultSearchCoins(sOpt)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   coins,
		"len":    len(coins),
	})
}

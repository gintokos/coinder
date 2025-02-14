package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/models"
	"github.com/gintokos/coinder/pkg/gerror"
	"github.com/gintokos/coinder/pkg/middleware"
)

type CoinService interface {
	DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.DBCoin, error)
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

	coingrp.Use(middleware.QueryMiddleware())

	coingrp.POST("/default", h.Default)
	// coingrp.GET("/custom", h.Custom)
}

// to do refactor on diff errors from service
func (h *CoinHandler) Default(c *gin.Context) {
	slog.Info("handlers.coin.Default")

	sOpt := middleware.SearchCoinOpt(c)
	coins, err := h.service.DefaultSearchCoins(sOpt)
	if err != nil {
		switch {
		case gerror.IsNotFound(err):
			c.JSON(http.StatusNotFound, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		case gerror.IsInternal(err):
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		}
		c.Set("error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   coins,
	})

}

// func (h *CoinHandler) Custom(c *gin.Context) {
// 	slog.Info("handlers.coin.Custom")

// 	user := UserFromClaims(c)

// 	var opt models.SearchCoinOpt
// 	err := c.BindJSON(&opt)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "error",
// 			"error":  "Invalid options data",
// 		})
// 		return
// 	}

// 	opt.UserID = user.ID
// 	coins, err := h.service.DefaultSearchCoins(opt)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "success",
// 		"data":   coins,
// 	})
// }

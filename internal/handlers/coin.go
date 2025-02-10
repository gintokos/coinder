package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/models"
)

const (
	BY_PRICE      = "BY_PRICE"
	BY_MARKET_CAP = "BY_MARKET_CAP"
)

func QueryOpt(c *gin.Context) models.QuerySearchCoinOpt {
	opt := c.MustGet("query")
	return opt.(models.QuerySearchCoinOpt)
}

func SearchCoinOpt(c *gin.Context) models.SearchCoinOpt {
	qOpt := QueryOpt(c)
	user := UserFromClaims(c)

	sOpt := models.SearchCoinOpt{
		UserID:             user.ID,
		QuerySearchCoinOpt: qOpt,
	}

	return sOpt
}

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

	coingrp.Use(QueryMiddleware)

	coingrp.POST("/default", h.Default)
	// coingrp.GET("/custom", h.Custom)
}

func QueryMiddleware(c *gin.Context) {
	sorteBy := c.DefaultQuery("sorteBy", "BY_MARKET_CAP")
	if sorteBy != BY_PRICE && sorteBy != BY_MARKET_CAP {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid sorteBy param",
		})
		return
	}

	page := c.DefaultQuery("page", "1")
	pageNum, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid page param",
		})
		return
	}

	limit := c.DefaultQuery("limit", "100")
	limitNum, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid limit param",
		})
		return
	}

	likedByUserStr := c.DefaultQuery("likedByUser", "false")
	likedByUser, err := strconv.ParseBool(likedByUserStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid likedByUser param",
		})
		return
	}

	if pageNum < 1 || limitNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "page and limit must be greater than 0",
		})
		return
	}

	c.Set("query", models.QuerySearchCoinOpt{
		Page:        int(pageNum),
		Limit:       int(limitNum),
		LikedByUser: likedByUser,
		SortedBy:    sorteBy,
	})

	c.Next()
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

// to do refactor on diff errors from service
func (h *CoinHandler) Default(c *gin.Context) {
	slog.Info("handlers.coin.Default")

	sOpt := SearchCoinOpt(c)
	coins, err := h.service.DefaultSearchCoins(sOpt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   coins,
	})

}

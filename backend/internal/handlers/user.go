package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/middleware"
	"github.com/gintokos/coinder/internal/models"
	"github.com/gintokos/coinder/pkg/sl"
)

type UserService interface {
	Update(models.User) error
	CreateInvoice(amount int) (string, error)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(router *gin.RouterGroup, service UserService) UserHandler {
	h := UserHandler{service}
	h.RegisterRoutes(router)
	return h
}

func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	usergrp := router.Group("/user")

	usergrp.GET("/:id", h.User)
	usergrp.POST("/update", h.Update)
	usergrp.GET("/create_invoice", h.CreateInvoice)

	slog.Info("UserHandler registered routes")
}

func (h *UserHandler) CreateInvoice(c *gin.Context) {
	amountstr := c.Query("amount")
	amount, err := strconv.Atoi(amountstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid amount parameter",
		})
		slog.Info("invalid amount parameter on create invoice", sl.Err(err))
		return
	}

	invoicestr, err := h.service.CreateInvoice(amount)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   invoicestr,
	})
}

func (h *UserHandler) User(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
	})
}

func (h *UserHandler) Update(c *gin.Context) {
	slog.Info("handlers.user.update")
	user := middleware.UserFromClaims(c)

	err := h.service.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
	})
}

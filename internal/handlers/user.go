package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/middleware"
	"github.com/gintokos/coinder/internal/models"
)

type UserService interface {
	Update(models.User) error
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

	slog.Info("UserHandler registered routes")
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

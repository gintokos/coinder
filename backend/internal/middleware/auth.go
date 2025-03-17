package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/backend/internal/models"
	"github.com/gintokos/coinder/backend/pkg/telegram"
)

func UserFromClaims(c *gin.Context) models.User {
	cl, _ := c.Get("claims")
	claims, ok := cl.(*telegram.TClaims)
	if !ok {
		panic("invalid claims type in context")
	}

	return models.User{
		ID:        claims.ID,
		FirstName: claims.FirstName,
		Username:  claims.Username,
		PhotoUrl:  claims.PhotoURL,
		AuthDate:  claims.AuthDate,
	}
}

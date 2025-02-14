package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/models"
	"github.com/golang-jwt/jwt"
)

func UserFromClaims(c *gin.Context) models.User {
	cl, _ := c.Get("claims")
	claims := *cl.(*jwt.MapClaims)

	return models.User{
		ID:        claims["id"].(int64),
		FirstName: claims["first_name"].(string),
		Username:  claims["username"].(string),
		PhotoUrl:  claims["photo_url"].(string),
		AuthDate:  claims["auth_date"].(string),
	}

}

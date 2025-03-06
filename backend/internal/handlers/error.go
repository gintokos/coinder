package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/pkg/gerror"
)

func ErrorResponse(c *gin.Context, err error) {
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
}

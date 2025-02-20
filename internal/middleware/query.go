package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/internal/constants"
	"github.com/gintokos/coinder/internal/models"
)

func sendBadReq(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"error":  message,
	})
	c.Abort()
}

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


func QueryMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		sorteBy := c.DefaultQuery("sorteBy", "BY_MARKET_CAP")
		if sorteBy != constants.BY_PRICE && sorteBy != constants.BY_MARKET_CAP {
			sendBadReq(c, "Invalid sorteBy param")
			return
		}

		page := c.DefaultQuery("page", "1")
		pageNum, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			sendBadReq(c, "Invalid page param")
			return
		}

		limit := c.DefaultQuery("limit", "100")
		limitNum, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			sendBadReq(c, "Invalid limit param")
			return
		}

		likedByUserStr := c.DefaultQuery("likedByUser", "true")
		likedByUser, err := strconv.ParseBool(likedByUserStr)
		if err != nil {
			sendBadReq(c, "Invalid likedByUser param")
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
}

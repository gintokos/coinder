package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/coinder/pkg/gerror"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		attrs := []any{
			"path", path,
			"status", status,
			"latency", latency.Microseconds(),
		}

		if e, ok := c.Get("error"); ok {
			gerr := e.(error)
			attrs = append(attrs, "error", gerror.FullError(gerr))
		}

		slog.Info("request processed", attrs...)
	}
}

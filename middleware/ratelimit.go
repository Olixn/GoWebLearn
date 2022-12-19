package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// RateLimitMiddleware 限流中间件
func RateLimitMiddleware(fillInterval time.Duration, capacity int64) func(ctx *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, capacity)
	return func(c *gin.Context) {
		// if bucket.Take(1) > 0 {
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

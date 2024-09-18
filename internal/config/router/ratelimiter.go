package router

import (
	"backend/pkg/handler"
	"net/http"
	"time"

	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-gonic/gin"
)

const REQUEST_LIMIT = 100
const LIMIT_INTERVAL = time.Second * 2

func rateLimiterConfig() gin.HandlerFunc {
	return ratelimiter.GinMemRatelimiter(ratelimiter.GinRatelimiterConfig{
		LimitKey: func(c *gin.Context) string {
			return c.ClientIP()
		},
		LimitedHandler: func(c *gin.Context) {
			handler.Error(c, http.StatusTooManyRequests, "Terlalu banyak permintaan")
			c.Abort()
		},
		TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
			return LIMIT_INTERVAL, REQUEST_LIMIT
		},
	})
}

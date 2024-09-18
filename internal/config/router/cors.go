package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsHeaderConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	}
}

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders: []string{
			"Content-Type", "refresh_token", "access_token", "uuid", "resetpw_token", "device",
		},
		ExposeHeaders: []string{
			"Content-Length", "refresh_token", "access_token", "uuid", "resetpw_token",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

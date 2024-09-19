package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsHeaderConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Origin") == "http://localhost:3000" || c.Request.Header.Get("Origin") == "http://127.0.0.1:3000" {
			c.Writer.Header().Add("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		}
		c.Next()
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

package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsHeaderConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if c.Request.Header.Get("Origin") == "http://localhost:5173" {
		// 	c.Writer.Header().Add("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		// }
		// c.Next()
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	}
}

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

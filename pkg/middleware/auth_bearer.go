package middleware

import (
	"backend/pkg/env"
	"backend/pkg/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthBearer() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")

		isValid := validateBearerToken(authToken)
		if !isValid {
			handler.Error(c, http.StatusUnauthorized, "Authorization header is required")
			return
		}

		if env.APIKey == "" {
			handler.Error(c, http.StatusInternalServerError, "API Key is not set")
			return
		}

		if authToken[7:] != env.APIKey {
			handler.Error(c, http.StatusUnauthorized, "Invalid Authorization header")
			return
		}

		c.Next()
	}
}

func validateBearerToken(token string) bool {
	if token == "" || len(token) < 7 || token[:7] != "Bearer " {
		return false
	}

	return true
}

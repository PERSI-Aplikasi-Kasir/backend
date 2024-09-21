package router

import (
	"backend/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var routerInstance *gin.Engine

func InitializeRouter() {
	fmt.Println("===== Initialize Router =====")
	router := gin.Default()
	router.Use(corsHeaderConfig())
	router.Use(corsConfig())
	router.Use(rateLimiterConfig())
	// TODO: discord logger not working in this microservice
	router.Use(logger.DiscordLogger())

	routerInstance = router

	fmt.Println("✓ Gin router initialized")
}

func GetRouterInstance() *gin.Engine {
	return routerInstance
}

func UnsyncRouter(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Error while shutting down the server")
		return
	}

	fmt.Println("✓ Router closed")
}

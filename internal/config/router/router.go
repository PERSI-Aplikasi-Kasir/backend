package router

import (
	"backend/database"
	userController "backend/internal/module/user/controller"
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
	fmt.Println("===== Inisialisasi Router =====")
	router := gin.Default()
	router.Use(corsHeaderConfig())
	router.Use(corsConfig())
	router.Use(rateLimiterConfig())
	router.Use(logger.DiscordLogger())

	routerInstance = router

	fmt.Println("✓ Gin router diinisialisasi")
}

func GetRouterInstance() *gin.Engine {
	return routerInstance
}

func InitializeRoutes() {
	fmt.Println("===== Inisialisasi Routes =====")
	router := GetRouterInstance()
	db := database.GetDBInstance()

	userController.NewUserController(db).Register(router)

	fmt.Println("✓ Inisialisasi", len(router.Routes()), "routes")
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

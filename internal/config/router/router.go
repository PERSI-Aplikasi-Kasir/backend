package router

import (
	"backend/database"
	userController "backend/internal/module/user/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

var routerInstance *gin.Engine

func InitializeRouter() {
	fmt.Println("===== Inisialisasi Router =====")
	router := gin.Default()
	router.Use(corsHeaderConfig())
	router.Use(corsConfig())
	router.Use(rateLimiterConfig())

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

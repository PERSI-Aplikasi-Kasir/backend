package router

import (
	whatsappController "backend/microservices/whatsappclient/chore/controller"
	"fmt"
)

func InitializeRoutes() {
	fmt.Println("===== Initialize Routes =====")
	router := GetRouterInstance()

	whatsappController := whatsappController.NewWhatsappController()
	whatsappController.Register(router)
	whatsappController.RegisterStream(router)

	fmt.Println("✓ Initialize", len(router.Routes()), "routes")
}

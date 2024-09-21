package router

import (
	whatsappController "backend/microservices/whatsappclient/chore/controller"
	"fmt"
)

func InitializeRoutes() {
	fmt.Println("===== Initialize Routes =====")
	router := GetRouterInstance()

	whatsappController.NewWhatsappController().Register(router)
	whatsappController.NewWhatsappController().RegisterStream(router)

	fmt.Println("✓ Initialize", len(router.Routes()), "routes")
}

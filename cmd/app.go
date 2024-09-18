package cmd

import (
	"backend/database"
	"backend/internal/config/router"
	"backend/internal/integration/mailer"
	"backend/pkg/env"
	"fmt"
)

func App() {
	fmt.Println("Running app...")

	database.InitializeDB()
	router.InitializeRouter()
	router.InitializeRoutes()
	mailer.InitializeMailer()

	routerInstance := router.GetRouterInstance()
	routerInstance.Run(
		env.BEHost + ":" + env.BEPort,
	)

	fmt.Println("✓ App is running")
}

package cmd

import (
	"backend/microservices/whatsappclient"
	"backend/pkg/env"
	"backend/pkg/logger"
	"fmt"
)

func WhatsappClient() {
	fmt.Println("Running microservice: Whatsapp Client")

	logger.InitializeLogger(env.LogsPath + "whatsappclient.log")
	whatsappclient.InitializeWhatsappClient()
}

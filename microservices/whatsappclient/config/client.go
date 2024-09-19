package config

import (
	"backend/microservices/whatsappclient/database"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.mau.fi/whatsmeow"
)

var client *whatsmeow.Client

func InitializeClient() {
	fmt.Println("===== Initialize Client =====")
	dbInstance := database.GetDBInstance()
	device, err := dbInstance.GetFirstDevice()
	if err != nil {
		log.Error().Err(err).Msg("failed to get first device from database")
		return
	}

	client = whatsmeow.NewClient(device, nil)

	if err := client.Connect(); err != nil {
		log.Error().Err(err).Msg("failed to connect to whatsapp")
		client.Disconnect()
		return
	}

	fmt.Println("✓ Client initialized")
}

func GetClient() *whatsmeow.Client {
	return client
}

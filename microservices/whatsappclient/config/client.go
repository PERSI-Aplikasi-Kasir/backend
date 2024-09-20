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

	if client.Store.ID != nil {
		if err := client.Connect(); err != nil {
			log.Error().Err(err).Msg("failed to connect to whatsapp")
			return
		}
	}

	fmt.Println("✓ Client initialized")
}

func GetClient() *whatsmeow.Client {
	if client == nil {
		InitializeClient()
	}

	return client
}

func ResyncClient(callerClient **whatsmeow.Client) error {
	if client.IsConnected() {
		client.Disconnect()
	}

	if client != nil {
		client = nil
	}

	InitializeClient()
	if client == nil {
		return fmt.Errorf("failed to reinitialize client")
	}

	*callerClient = client

	if client.Store.ID != nil {
		if err := client.Connect(); err != nil {
			log.Error().Err(err).Msg("failed to connect to whatsapp")
			return err
		}
	}
	return nil
}

func UnsyncClient() {
	if client != nil {
		client.Disconnect()
		client = nil
		fmt.Println("✓ Client closed")
	}
}

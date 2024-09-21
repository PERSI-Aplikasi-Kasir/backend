package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var (
	ENVIRONMENT string

	// server
	BEHost string
	BEPort string
	FEHost string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Services related
	APIKey            string
	UserAdminEmail    string
	UserAdminPassword string
	MailerEmail       string
	MailerPassword    string
	ResetPWFEEndpoint string
	DiscordWebhookUrl string
	LogsPath          string

	// Microservices
	LoggerPort     string
	WAClientPort   string
	WAClientDBName string
)

func InitializeEnv() {
	fmt.Println("===== Initialize .env =====")

	err := godotenv.Load(".env")
	if err != nil {
		log.Error().Err(err).Msg("Error while loading .env file")
		panic(err)
	}

	ENVIRONMENT = os.Getenv("ENVIRONMENT")
	BEHost = os.Getenv("BE_HOST")
	BEPort = os.Getenv("BE_PORT")
	FEHost = os.Getenv("FE_HOST")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	APIKey = os.Getenv("API_KEY")
	UserAdminEmail = os.Getenv("USER_ADMIN_EMAIL")
	UserAdminPassword = os.Getenv("USER_ADMIN_PASSWORD")
	MailerEmail = os.Getenv("MAILER_EMAIL")
	MailerPassword = os.Getenv("MAILER_PASSWORD")
	ResetPWFEEndpoint = os.Getenv("RESETPW_FE_ENDPOINT")
	DiscordWebhookUrl = os.Getenv("DISCORD_WEBHOOK_URL")
	LogsPath = os.Getenv("LOGS_PATH")
	LoggerPort = os.Getenv("LOGGER_PORT")
	WAClientPort = os.Getenv("WACLIENT_PORT")
	WAClientDBName = os.Getenv("WACLIENT_DB_NAME")

	fmt.Println("✓ .env initialized")
}

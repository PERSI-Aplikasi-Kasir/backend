package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	ENVIRONMENT       string
	BEHost            string
	BEPort            string
	FEHost            string
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	JWTSecretKey      string
	UserAdminEmail    string
	UserAdminPassword string
	MailerEmail       string
	MailerPassword    string
	ResetPWFEEndpoint string
	DiscordWebhookUrl string
)

func InitializeEnv() {
	fmt.Println("===== Initialize .env =====")

	err := godotenv.Load(".env")
	if err != nil {
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
	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	UserAdminEmail = os.Getenv("USER_ADMIN_EMAIL")
	UserAdminPassword = os.Getenv("USER_ADMIN_PASSWORD")
	MailerEmail = os.Getenv("MAILER_EMAIL")
	MailerPassword = os.Getenv("MAILER_PASSWORD")
	ResetPWFEEndpoint = os.Getenv("RESETPW_FE_ENDPOINT")
	DiscordWebhookUrl = os.Getenv("DISCORD_WEBHOOK_URL")

	fmt.Println("✓ .env initialized")
}

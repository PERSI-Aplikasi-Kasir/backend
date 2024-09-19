package mailer

import (
	"backend/pkg/env"
	"fmt"
	"net/smtp"
)

type Config struct {
	Auth   smtp.Auth
	From   string
	Server string
}

var mailerInstance *Config

func InitializeMailer() {
	fmt.Println("===== Initialize Mailer =====")

	email := env.MailerEmail
	password := env.MailerPassword

	mailerInstance = &Config{
		Auth:   smtp.PlainAuth("", email, password, "smtp.gmail.com"),
		From:   email,
		Server: "smtp.gmail.com:587",
	}

	fmt.Println("✓ Mailer initialized")
}

func GetMailerInstance() *Config {
	return mailerInstance
}

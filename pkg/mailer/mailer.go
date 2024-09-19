package mailer

import (
	"backend/internal/integration/mailer"
	"fmt"
	"net/smtp"

	"github.com/rs/zerolog/log"
)

type MailInfo struct {
	EmailTarget []string
	Subject     string
	Body        string
}

func SendMail(mailInfo MailInfo) {
	go func() {
		mailerInstance := mailer.GetMailerInstance()
		message := []byte(fmt.Sprintf("Subject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s\r\n", mailInfo.Subject, mailInfo.Body))

		err := smtp.SendMail(mailerInstance.Server, mailerInstance.Auth, mailerInstance.From, mailInfo.EmailTarget, message)
		if err != nil {
			log.Error().Err(err).Msgf("Error sending email to: %v", mailInfo.EmailTarget)
			return
		}
	}()
}

package utils

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"gopkg.in/gomail.v2"
)

var d = func() *gomail.Dialer {
	smtpEmail := os.Getenv("SMTP_SERVER")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Error(err)
	}
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPass := os.Getenv("SMTP_PASSWORD")
	return gomail.NewDialer(smtpEmail, smtpPort, smtpUsername, smtpPass)
}()

func SendEmail(email string, subject string, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("SMTP_USERNAME"))
	m.SetHeader("")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

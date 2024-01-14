package commands

import (
	"net/smtp"

	config "github.com/datrine/alumni_business/Config"
	"github.com/matcornic/hermes/v2"
)

type SendEmailData struct {
	Email   string
	Message *hermes.Body
}

func SendEmail(data *SendEmailData) error {
	from := config.GetSMTPUsername()
	username := config.GetSMTPUsername()
	password := config.GetSMTPPassword()
	host := config.GetSMTPUsername()
	auth := smtp.PlainAuth("", username, password, host)
	to := []string{data.Email}
	err := smtp.SendMail(data.Email, auth, from, to, []byte{})
	if err != nil {
		return err
	}

	return nil
}

func SendEmailHermes(data *SendEmailData) error {
	h := hermes.Hermes{}
	email := hermes.Email{
		Body: *data.Message,
	}
	em, _ := h.GenerateHTML(email)
	from := config.GetSMTPUsername()
	username := config.GetSMTPUsername()
	password := config.GetSMTPPassword()
	host := config.GetSMTPHost()
	port := config.GetSMTPPort()
	auth := smtp.PlainAuth("", username, password, host)
	to := []string{data.Email}
	err := smtp.SendMail(host+":"+port, auth, from, to, []byte(em))
	if err != nil {
		return err
	}

	return nil
}

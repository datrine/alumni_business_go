package utils

import (
	"fmt"

	config "github.com/datrine/alumni_business/Config"
	"github.com/matcornic/hermes/v2"
	"github.com/resend/resend-go/v2"
)

type SendEmailData struct {
	Email   string
	Subject string
	Message *hermes.Body
}

func SendEmailHermes(data *SendEmailData) error {
	h := hermes.Hermes{}
	email := hermes.Email{
		Body: *data.Message,
	}
	em, _ := h.GenerateHTML(email)
	/*from := config.GetSMTPUsername() //"onboarding@resend.dev"
	username := config.GetSMTPUsername()
	password := config.GetSMTPPassword()
	host := config.GetSMTPHost()
	port := config.GetSMTPPort()
	auth := smtp.PlainAuth("", username, password, host)
	to := []string{data.Email}
	fmt.Println(em)
	err := smtp.SendMail(host+":"+port, auth, from, to, []byte(em))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
	*/
	apiKey := config.GetResendAPIKey()

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    config.GetResendFromEmail(),
		To:      []string{data.Email},
		Subject: data.Subject,
		Html:    em,
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(sent)
	return nil
}

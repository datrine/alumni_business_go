package paystack

import (
	"encoding/json"
	"errors"
	"strings"

	config "github.com/datrine/alumni_business/Config"
	"github.com/gofiber/fiber/v2"
)

func GeneratePaymentLink(data *PaystackTransactionRequestJSON) (*PaystackTransactionResponseJSON, error) {
	paystackSK := config.GetPaystackSecretKey()
	agent := fiber.Post("https://api.paystack.co/transaction/initialize")
	agent.Set("Authorization", "Bearer "+paystackSK)
	agent.Set("Content-Type", "application/json")
	agent.JSON(data)
	status, dataInBytes, err := agent.Bytes()
	if len(err) > 0 {
		return nil, err[0]
	}
	if status >= 200 && status <= 299 {
		rrr := &PaystackTransactionResponseJSON{}
		err := json.Unmarshal(dataInBytes, rrr)
		return rrr, err
	}
	println(status)
	println(string(dataInBytes))
	return nil, errors.New("failed")
}

func VerifyPayment(reference string) (*PaystackVerifyTransactionResponseJSON, error) {
	paystackSK := config.GetPaystackSecretKey()
	agent := fiber.Get(strings.Join([]string{"https://api.paystack.co/transaction/verify", reference}, "/"))
	agent.Set("Authorization", "Bearer "+paystackSK)
	status, dataInBytes, err := agent.Bytes()
	if len(err) > 0 {
		return nil, err[0]
	}
	if status >= 200 && status <= 299 {
		rrr := &PaystackVerifyTransactionResponseJSON{}
		err := json.Unmarshal(dataInBytes, rrr)
		return rrr, err
	}
	return nil, errors.New("something when wrong")
}

type PaystackVerifyTransactionRequestJSON struct {
	Email  string
	Amount string
}

type PaystackTransactionRequestJSON struct {
	Email  string `json:"email,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type PaystackTransactionResponseJSON struct {
	Status bool                    `json:"status"`
	Amount string                  `json:"amount"`
	Data   PaystackTransactionData `json:"data"`
}

type PaystackVerifyTransactionResponseJSON struct {
	Status bool                    `json:"status"`
	Amount string                  `json:"amount"`
	Data   PaystackTransactionData `json:"data"`
}

type PaystackTransactionData struct {
	AuthorizationUrl string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

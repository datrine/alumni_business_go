package paystack

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	config "github.com/datrine/alumni_business/Config"
	models "github.com/datrine/alumni_business/Models"
	providers "github.com/datrine/alumni_business/Providers"
	"github.com/gofiber/fiber/v2"
)

func GeneratePaymentLink(data *GeneratePaymentLinkDTO) (*PaystackTransactionResponseJSON, error) {
	paystackSK := config.GetPaystackSecretKey()
	agent := fiber.Post("https://api.paystack.co/transaction/initialize")
	agent.Set("Authorization", "Bearer "+paystackSK)
	agent.Set("Content-Type", "application/json")
	requestData := &PaystackTransactionRequestJSON{
		Email:  data.PayerEmail,
		Amount: data.Amount,
	}
	agent.JSON(requestData)
	status, dataInBytes, err := agent.Bytes()
	if len(err) > 0 {
		return nil, err[0]
	}
	if status >= 200 && status <= 299 {
		rrr := &PaystackTransactionResponseJSON{}
		err := json.Unmarshal(dataInBytes, rrr)
		if err != nil {
			fmt.Printf(err.Error())
		}
		_, err = CreateTransaction(&models.Transaction{
			PayerID:    data.PayerID,
			PayerEmail: data.PayerEmail,
			Amount:     data.Amount,
			ID:         rrr.Data.Reference,
			Metadata:   rrr,
			Status:     "INITIALIZED",
			Currency:   "NGN",
			Platform:   "PAYSTACK",
		})
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

func CreateTransaction(data *models.Transaction) (*models.Transaction, error) {
	result := providers.DB.Create(data)
	err := result.Error
	return data, err
}

type PaystackVerifyTransactionRequestJSON struct {
	Email  string
	Amount string
}

type GeneratePaymentLinkDTO struct {
	Amount     string
	PayerID    string
	Reference  string
	PayerEmail string
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

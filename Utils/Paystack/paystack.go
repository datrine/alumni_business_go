package paystack

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

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

func VerifyPaymentsJob() {
	unverifiedPayments := &[]models.Transaction{}
	fmt.Println("Running")
	result := providers.DB.Where(&models.Transaction{
		Status: "INITIALIZED",
	}).Find(unverifiedPayments)
	err := result.Error
	if err != nil {
		return
	}
	for _, tx := range *unverifiedPayments {
		VerifyPayment(tx.ID)
	}
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
		if rrr.Data.Status == "success" {
			transactionModel := models.Transaction{
				ID: reference,
			}
			providers.DB.Model(transactionModel).UpdateColumn("status", "VERIFIED")
		}
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

type PaystackTransactionData struct {
	AuthorizationUrl string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type LogHistory struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

type Log struct {
	StartTime int           `json:"start_time"`
	TimeSpent int           `json:"time_spent"`
	Attempts  int           `json:"attempts"`
	Errors    int           `json:"errors"`
	Success   bool          `json:"success"`
	Mobile    bool          `json:"mobile"`
	Input     []interface{} `json:"input"`
	History   []LogHistory  `json:"history"`
}

type Authorization struct {
	AuthorizationCode string      `json:"authorization_code"`
	Bin               string      `json:"bin"`
	Last4             string      `json:"last4"`
	ExpMonth          string      `json:"exp_month"`
	ExpYear           string      `json:"exp_year"`
	Channel           string      `json:"channel"`
	CardType          string      `json:"card_type"`
	Bank              string      `json:"bank"`
	CountryCode       string      `json:"country_code"`
	Brand             string      `json:"brand"`
	Reusable          bool        `json:"reusable"`
	Signature         string      `json:"signature"`
	AccountName       interface{} `json:"account_name"`
}

type Customer struct {
	ID                       int         `json:"id"`
	FirstName                interface{} `json:"first_name"`
	LastName                 interface{} `json:"last_name"`
	Email                    string      `json:"email"`
	CustomerCode             string      `json:"customer_code"`
	Phone                    interface{} `json:"phone"`
	Metadata                 interface{} `json:"metadata"`
	RiskAction               string      `json:"risk_action"`
	InternationalFormatPhone interface{} `json:"international_format_phone"`
}

type Data struct {
	ID                 int           `json:"id"`
	Domain             string        `json:"domain"`
	Status             string        `json:"status"`
	Reference          string        `json:"reference"`
	Amount             int           `json:"amount"`
	Message            interface{}   `json:"message"`
	GatewayResponse    string        `json:"gateway_response"`
	PaidAt             time.Time     `json:"paid_at"`
	CreatedAt          time.Time     `json:"created_at"`
	Channel            string        `json:"channel"`
	Currency           string        `json:"currency"`
	IPAddress          string        `json:"ip_address"`
	Metadata           string        `json:"metadata"`
	Log                Log           `json:"log"`
	Fees               int           `json:"fees"`
	FeesSplit          interface{}   `json:"fees_split"`
	Authorization      Authorization `json:"authorization"`
	Customer           Customer      `json:"customer"`
	Plan               interface{}   `json:"plan"`
	Split              struct{}      `json:"split"`
	OrderID            interface{}   `json:"order_id"`
	PaidAtAlt          time.Time     `json:"paidAt"`
	CreatedAtAlt       time.Time     `json:"createdAt"`
	RequestedAmount    int           `json:"requested_amount"`
	POSTransactionData interface{}   `json:"pos_transaction_data"`
	Source             interface{}   `json:"source"`
	FeesBreakdown      interface{}   `json:"fees_breakdown"`
	TransactionDate    time.Time     `json:"transaction_date"`
	PlanObject         struct{}      `json:"plan_object"`
	Subaccount         struct{}      `json:"subaccount"`
}

type PaystackVerifyTransactionResponseJSON struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

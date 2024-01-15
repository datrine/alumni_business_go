package config

import (
	"os"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func GetDBURL() string {
	return os.Getenv("DB_URL")
}

func GetSMTPUsername() string {
	return os.Getenv("SMTP_USERNAME")
}

func GetSMTPPassword() string {
	return os.Getenv("SMTP_PASSWORD")
}

func GetSMTPHost() string {
	return os.Getenv("SMTP_HOST")
}

func GetSMTPPort() string {
	return os.Getenv("SMTP_PORT")
}

func GetResendAPIKey() string {
	return os.Getenv("RESEND_API_KEY")
}

func GetResendFromEmail() string {
	return os.Getenv("RESEND_FROM_EMAIL")
}

func GetPaystackSecretKey() string {
	return os.Getenv("PAYSTACK_SECRET_KEY")
}

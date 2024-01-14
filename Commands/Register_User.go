package commands

import (
	"fmt"

	dtos "github.com/datrine/alumni_business/Dtos/Command"
	entities "github.com/datrine/alumni_business/Entities"
	paystack "github.com/datrine/alumni_business/Utils/Paystack"
	"github.com/jaevor/go-nanoid"
	"github.com/matcornic/hermes/v2"
)

type UserEntityWithPaystackLink struct {
	User     *entities.User
	Paystack *paystack.PaystackTransactionResponseJSON
}

func RegisterUser(data *dtos.RegisterUserCommandDTO) (*UserEntityWithPaystackLink, error) {
	passwordGenerator, err := nanoid.ASCII(10)
	if err != nil {
		return nil, err
	}
	password := passwordGenerator()
	user := entities.User{
		Email:             data.Email,
		FirstName:         data.FirstName,
		LastName:          data.LastName,
		MemberNumber:      data.MemberNumber,
		ProfilePictureUrl: data.ProfilePictureUrl,
		Password:          password,
	}
	err = user.Register()
	if err != nil {
		return nil, err
	}
	paystackJSON, err := paystack.GeneratePaymentLink(&paystack.GeneratePaymentLinkDTO{
		PayerEmail: user.Email,
		Amount:     "60000",
		PayerID:    user.ID,
	})

	if err != nil {
		return nil, err
	}
	err = SendEmailHermes(&SendEmailData{Email: data.Email, Message: &hermes.Body{
		Name:     data.FirstName,
		Greeting: "Hi",
		Intros:   []string{"Thank you for your registration"},
		Actions: []hermes.Action{
			{
				Instructions: "Click on the click to pay for your membership",
				Button: hermes.Button{
					Link: paystackJSON.Data.AuthorizationUrl,
				},
			},
		},
	}})

	if err != nil {
		fmt.Println(err)
	}

	return &UserEntityWithPaystackLink{
		User:     &user,
		Paystack: paystackJSON,
	}, nil
}

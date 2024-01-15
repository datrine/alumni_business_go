package handlers

import (
	entities "github.com/datrine/alumni_business/Entities"
	paystack "github.com/datrine/alumni_business/Utils/Paystack"
	"github.com/gofiber/fiber/v2"
)

//var validate = validator.New(validator.WithRequiredStructEnabled())

type GeneratePaystackLinkErrorResponse struct {
	Status  int
	Message string
}

// Register: Generate payment link
//
//	@Summary      login user
//	@Description  login user
//	@Tags         accounts
//	@Accept       json
//	@Param        email query string false "email of registered user"
//	@Produce      json
//	@Success      200  {object}  UpdateUserProfileSuccessResponse
//	@Failure      400  {object}  UpdateUserProfileErrorResponse
//	@Failure      404  {object}  UpdateUserProfileErrorResponse
//	@Failure      500  {object}  UpdateUserProfileErrorResponse
//	@Router /generate_payment_link [get]
func GeneratePaystackLink(c *fiber.Ctx) error {
	emailQuery := c.Query("email")
	user, err := entities.GetUserByEmail(emailQuery)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&GeneratePaystackLinkErrorResponse{
			Status:  fiber.StatusNotFound,
			Message: err.Error(),
		})
	}
	genRes, err := paystack.GeneratePaymentLink(&paystack.GeneratePaymentLinkDTO{
		Amount:     "60000",
		PayerID:    user.ID,
		PayerEmail: user.Email,
	})
	return c.Status(fiber.StatusOK).JSON(&GeneratePaystackLinkSuccessResponse{
		Status:  fiber.StatusOK,
		Message: "Payment link generated successfully",
		Data: &GeneratePaystackLinkSuccessData{
			ID:                user.ID,
			Email:             user.Email,
			FirstName:         user.FirstName,
			LastName:          user.LastName,
			ProfilePictureUrl: user.ProfilePictureUrl,
			PaymentLink:       genRes.Data.AuthorizationUrl},
	})
}

type GeneratePaystackLinkSuccessData struct {
	ID                string `json:"id"`
	Token             string `json:"access_token"`
	Email             string `json:"email"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	PaymentLink       string `json:"payment_link"`
}

type GeneratePaystackLinkSuccessResponse struct {
	Message string                           `json:"message"`
	Status  int                              `json:"status"`
	Data    *GeneratePaystackLinkSuccessData `json:"data"`
}

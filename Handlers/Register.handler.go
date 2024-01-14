package handlers

import (
	"strings"

	commands "github.com/datrine/alumni_business/Commands"
	dtosCommand "github.com/datrine/alumni_business/Dtos/Command"
	dtosRequest "github.com/datrine/alumni_business/Dtos/Request"
	utils "github.com/datrine/alumni_business/Utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type RegisterUserErrorResponse struct {
	Status  int
	Message string
}

// Register: Register new user
//
//	@Summary      Register user
//	@Description  register user
//	@Tags         accounts
//	@Accept       mpfd
//	@Param        member_number formData string true "member number" minlength(5)  maxlength(10)
//	@Param        email formData string  true "email"
//	@Param        first_name formData string   true "first name"
//	@Param        last_name formData string   true "last name"
//	@Param	   	  profile_picture formData file true "profile picture"
//	@Produce      json
//	@Success      200  {object}  RegisterUserSuccessResponse
//	@Failure      400  {object}  RegisterUserErrorResponse
//	@Failure      404  {object}  RegisterUserErrorResponse
//	@Failure      500  {object}  RegisterUserErrorResponse
//	@Router /users/register [post]
func Register(c *fiber.Ctx) error {
	header := &c.Request().Header
	if strings.Contains(strings.ToLower(string(header.ContentType())), "multipart/form-data") {
		memberNumber := c.FormValue("member_number")
		firstName := c.FormValue("first_name")
		lastName := c.FormValue("last_name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		profilePictureFile, err := c.FormFile("profile_picture")
		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(&RegisterUserErrorResponse{
				Status:  fiber.ErrBadRequest.Code,
				Message: err.Error(),
			})
		}
		if _, err = profilePictureFile.Open(); err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(&RegisterUserErrorResponse{
				Status:  fiber.ErrBadRequest.Code,
				Message: err.Error(),
			})
		}
		profilePictureUrl, err := utils.UploadFile(profilePictureFile)
		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(RegisterUserErrorResponse{
				Status:  fiber.ErrBadRequest.Code,
				Message: err.Error(),
			})
		}
		data := &dtosRequest.RegisterUserRequestFORMDTO{
			MemberNumber:      memberNumber,
			Email:             email,
			FirstName:         firstName,
			LastName:          lastName,
			ProfilePictureUrl: profilePictureUrl,
			Password:          password,
		}
		err = validate.Struct(data)
		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(&RegisterUserErrorResponse{
				Message: err.Error(),
				Status:  fiber.ErrBadGateway.Code,
			})
		}

		entityUser, err := commands.RegisterUser(&dtosCommand.RegisterUserCommandDTO{
			MemberNumber:      data.MemberNumber,
			Email:             data.Email,
			FirstName:         data.FirstName,
			LastName:          data.LastName,
			ProfilePictureUrl: data.ProfilePictureUrl,
		})

		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(&utils.DefaultErrorResponse{
				Message: err.Error(),
				Status:  fiber.ErrBadGateway.Code,
			})
		}
		return c.Status(fiber.StatusCreated).JSON(&RegisterUserSuccessResponse{
			Message: "user registered successfully",
			Code:    fiber.StatusCreated,
			Data: &RegisterUserResponseData{
				ID:                entityUser.User.ID,
				FirstName:         entityUser.User.FirstName,
				LastName:          entityUser.User.LastName,
				Email:             entityUser.User.Email,
				ProfilePictureUrl: entityUser.User.ProfilePictureUrl,
				PaymentLink:       entityUser.Paystack.Data.AuthorizationUrl,
			},
		})
	}
	return nil
}

type RegisterUserResponseData struct {
	ID                string `json:"id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	PaymentLink       string `json:"payment_link"`
}

type RegisterUserSuccessResponse struct {
	Message string                    `json:"message"`
	Code    int                       `json:"status"`
	Data    *RegisterUserResponseData `json:"data"`
}

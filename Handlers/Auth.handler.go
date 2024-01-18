package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	commands "github.com/datrine/alumni_business/Commands"
	dtosCommand "github.com/datrine/alumni_business/Dtos/Command"
	dtosRequest "github.com/datrine/alumni_business/Dtos/Request"
	queries "github.com/datrine/alumni_business/Queries"
	utils "github.com/datrine/alumni_business/Utils"
	"github.com/gofiber/fiber/v2"
)

//var validate = validator.New(validator.WithRequiredStructEnabled())

type LoginErrorResponse struct {
	Status  int
	Message string
}

// Register: Update logged-in user
//
//	@Summary      login user
//	@Description  login user
//	@Tags         accounts
//	@Accept       json
//	@Accept       mpfd
//	@Param        identifier formData string   false "identifier: email or member number"
//	@Param        password formData string   false "password"
//	@Param        jsonData body dtosRequest.BasicLoginRequestJSONDTO   false "date of birth"
//	@Produce      json
//	@Success      200  {object}  UpdateUserProfileSuccessResponse
//	@Failure      400  {object}  UpdateUserProfileErrorResponse
//	@Failure      404  {object}  UpdateUserProfileErrorResponse
//	@Failure      500  {object}  UpdateUserProfileErrorResponse
//	@Router /login/basic [post]
func Login(c *fiber.Ctx) error {
	header := &c.Request().Header
	if strings.Contains(strings.ToLower(string(header.ContentType())), "application/json") {
		loginDetailsJSON := dtosRequest.BasicLoginRequestJSONDTO{}
		err := json.Unmarshal(c.Body(), &loginDetailsJSON)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&utils.DefaultErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
			})
		}
		fmt.Println(loginDetailsJSON)
		data := &dtosCommand.BasicLoginCommandDTO{
			Identifier: loginDetailsJSON.Identifier,
			Password:   loginDetailsJSON.Password,
		}
		err = validate.Struct(data)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&RegisterUserErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
			})
		}

		authEntityUser, err := commands.BasicLogin(&dtosCommand.BasicLoginCommandDTO{
			Identifier: data.Identifier,
			Password:   data.Password,
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(&utils.DefaultErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusUnauthorized,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&BasicLoginResponseDataSuccessResponse{
			Message: "Login successfully",
			Code:    fiber.StatusCreated,
			Data: &BasicLoginResponseData{
				ID:                authEntityUser.User.ID,
				Token:             authEntityUser.Token,
				FirstName:         authEntityUser.User.FirstName,
				LastName:          authEntityUser.User.LastName,
				ProfilePictureUrl: authEntityUser.User.ProfilePictureUrl,
			},
		})
	}
	return errors.New("wrong header field")
}

// Change password
//
//	@Summary      change password
//	@Description  change password
//	@Tags         accounts
//	@Accept       json
//	@Accept       mpfd
//	@Param        jsonChangePaywordData  body  dtosRequest.ChangePasswordRequestJSONDTO   false   "uu"
//	@Produce      json
//	@Success      200  {object}  ChangePasswordSuccessResponse
//	@Failure      400  {object}  UpdateUserProfileErrorResponse
//	@Failure      404  {object}  UpdateUserProfileErrorResponse
//	@Failure      500  {object}  UpdateUserProfileErrorResponse
//	@Router /auth/password/change [post]
func ChangePassword(c *fiber.Ctx) error {
	header := &c.Request().Header
	if strings.Contains(strings.ToLower(string(header.ContentType())), "application/json") {
		changePasswordJSON := dtosRequest.ChangePasswordRequestJSONDTO{}
		err := json.Unmarshal(c.Body(), &changePasswordJSON)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&utils.DefaultErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
			})
		}
		fmt.Println(changePasswordJSON)
		data := &dtosCommand.ChangePasswordCommandDTO{
			OldPassword: changePasswordJSON.OldPassword,
			NewPassword: changePasswordJSON.NewPassword,
		}
		err = validate.Struct(data)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&RegisterUserErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
			})
		}
		payload, err := utils.GetAuthPayload(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&RegisterUserErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
			})
		}
		authUser, err := commands.UpdateUserPassword(&dtosCommand.UpdateUserPasswordCommandDTO{
			OldPassword: data.OldPassword,
			NewPassword: data.NewPassword,
			Email:       payload.Email,
			ID:          payload.ID,
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(&utils.DefaultErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusUnauthorized,
			})
		}
		return c.Status(fiber.StatusOK).JSON(&ChangePasswordSuccessResponse{
			Message: "Password changed successfully",
			Code:    fiber.StatusOK,
			Data: &ChangePasswordResponseData{
				Token:     authUser.Token,
				ID:        authUser.User.ID,
				FirstName: authUser.User.FirstName,
			},
		})
	}
	return errors.New("wrong header field")
}

// @Summary      Get my profile
// @Description  Get my profile
// @Tags         accounts
// @Produce      json
// @Success      200  {object}  ChangePasswordSuccessResponse
// @Failure      400  {object}  UpdateUserProfileErrorResponse
// @Failure      404  {object}  UpdateUserProfileErrorResponse
// @Failure      500  {object}  UpdateUserProfileErrorResponse
// @Router /auth/me [get]
func GetMyProfile(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	authHeaderValue := headers["Authorization"][0]
	tokenString := strings.Split(authHeaderValue, " ")[1]
	payload, err := utils.GetAuthPayload(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&RegisterUserErrorResponse{
			Message: err.Error(),
			Status:  fiber.StatusBadRequest,
		})
	}
	authUser, err := queries.GetAuthUserByEmail(payload.Email)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&utils.DefaultErrorResponse{
			Message: err.Error(),
			Status:  fiber.StatusUnauthorized,
		})
	}
	return c.Status(fiber.StatusOK).JSON(&GetMyProfileSuccessResponse{
		Message: "My profile",
		Code:    fiber.StatusOK,
		Data: &GetMyProfileSuccessResponseData{
			Token:             tokenString,
			ID:                authUser.ID,
			FirstName:         authUser.FirstName,
			LastName:          authUser.LastName,
			ProfilePictureUrl: authUser.ProfilePictureUrl,
			DOB:               authUser.DOB,
		},
	})
}

type BasicLoginResponseDataSuccessResponse struct {
	Message string                  `json:"message"`
	Code    int                     `json:"status"`
	Data    *BasicLoginResponseData `json:"data"`
}

type BasicLoginResponseData struct {
	ID                string `json:"id"`
	Token             string `json:"access_token"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	ProfilePictureUrl string `json:"profile_picture_url"`
}

type ChangePasswordSuccessResponse struct {
	Message string                      `json:"message"`
	Code    int                         `json:"status"`
	Data    *ChangePasswordResponseData `json:"data"`
}

type ChangePasswordResponseData struct {
	ID                string `json:"id"`
	Token             string `json:"access_token"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	ProfilePictureUrl string `json:"profile_picture_url"`
}

type GetMyProfileSuccessResponse struct {
	Message string                           `json:"message"`
	Code    int                              `json:"status"`
	Data    *GetMyProfileSuccessResponseData `json:"data"`
}

type GetMyProfileSuccessResponseData struct {
	ID                string     `json:"id"`
	Token             string     `json:"access_token"`
	FirstName         string     `json:"first_name"`
	Email             string     `json:"email"`
	LastName          string     `json:"last_name"`
	ProfilePictureUrl string     `json:"profile_picture_url"`
	DOB               *time.Time `json:"dob"`
}

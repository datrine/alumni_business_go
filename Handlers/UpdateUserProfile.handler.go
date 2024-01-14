package handlers

import (
	"fmt"
	"strings"
	"time"

	commands "github.com/datrine/alumni_business/Commands"
	dtosCommand "github.com/datrine/alumni_business/Dtos/Command"
	dtosRequest "github.com/datrine/alumni_business/Dtos/Request"
	utils "github.com/datrine/alumni_business/Utils"
	"github.com/gofiber/fiber/v2"
)

//var validate = validator.New(validator.WithRequiredStructEnabled())

type UpdateUserProfileErrorResponse struct {
	Status  int
	Message string
}

// Register: Update logged-in user
//
//	@Summary      Update user
//	@Description  Update user
//	@Tags         accounts
//	@Accept       mpfd
//	@Param        first_name formData string   false "first name"
//	@Param        last_name formData string   false "last name"
//	@Param        dob formData string   false "date of birth"
//	@Param	   	  profile_picture formData file false "profile picture"
//	@Produce      json
//	@Success      200  {object}  UpdateUserProfileSuccessResponse
//	@Failure      400  {object}  UpdateUserProfileErrorResponse
//	@Failure      404  {object}  UpdateUserProfileErrorResponse
//	@Failure      500  {object}  UpdateUserProfileErrorResponse
//	@Router /me/edit [post]
func UpdateUserProfile(c *fiber.Ctx) error {
	header := &c.Request().Header
	if strings.Contains(strings.ToLower(string(header.ContentType())), "multipart/form-data") {

		firstName := c.FormValue("first_name")
		lastName := c.FormValue("last_name")
		profession := c.FormValue("profession")
		dobString := c.FormValue("dob")
		profilePictureFile, err := c.FormFile("profile_picture")
		if err != nil {
			fmt.Printf(err.Error())
		}
		if _, err = profilePictureFile.Open(); err != nil {

			fmt.Printf(err.Error())
		}
		profilePictureUrl, err := utils.UploadFile(profilePictureFile)
		if err != nil {
			fmt.Printf(err.Error())
		}
		dob, err := time.Parse("", dobString)
		if err != nil {
			fmt.Printf(err.Error())
		}
		data := &dtosRequest.UpdateUserProfileRequestFORMDTO{
			FirstName:         firstName,
			LastName:          lastName,
			ProfilePictureUrl: profilePictureUrl,
			Profession:        &profession,
			DOB:               &dob,
		}
		err = validate.Struct(data)
		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(&RegisterUserErrorResponse{
				Message: err.Error(),
				Status:  fiber.ErrBadGateway.Code,
			})
		}

		entityUser, err := commands.UpdateUserProfile(&dtosCommand.UpdateUserProfileCommandData{
			FirstName:         data.FirstName,
			LastName:          data.LastName,
			Profession:        &profession,
			ProfilePictureUrl: data.ProfilePictureUrl,
		})

		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(&utils.DefaultErrorResponse{
				Message: err.Error(),
				Status:  fiber.ErrBadGateway.Code,
			})
		}
		return c.Status(fiber.StatusCreated).JSON(&UpdateUserProfileSuccessResponse{
			Message: "Profile updated successfully",
			Code:    fiber.StatusCreated,
			Data: &UpdateUserProfileResponseData{
				ID:                entityUser.ID,
				FirstName:         entityUser.FirstName,
				LastName:          entityUser.LastName,
				ProfilePictureUrl: entityUser.ProfilePictureUrl,
			},
		})
	}
	return nil
}

type UpdateUserProfileResponseData struct {
	ID                string `json:"id"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	ProfilePictureUrl string `json:"profile_picture_url,omitempty"`
}

type UpdateUserProfileSuccessResponse struct {
	Message string                         `json:"message"`
	Code    int                            `json:"status"`
	Data    *UpdateUserProfileResponseData `json:"data"`
}

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
	"github.com/golang-jwt/jwt/v5"
)

//var validate = validator.New(validator.WithRequiredStructEnabled())

type UpdateUserProfileErrorResponse struct {
	Status  int
	Message string
}

// Register: Update logged-in user
//
//	@Summary      logged-in user
//	@Description  logged-in user
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
//	@Router /auth/me/edit [post]
func UpdateUserProfile(c *fiber.Ctx) error {
	header := &c.Request().Header
	if strings.Contains(strings.ToLower(string(header.ContentType())), "multipart/form-data") {

		firstName := c.FormValue("first_name")
		lastName := c.FormValue("last_name")
		profession := c.FormValue("profession")
		dobString := c.FormValue("dob")
		profilePictureFile, err := c.FormFile("profile_picture")
		var profilePictureUrl string = ""
		if err != nil {
			fmt.Printf(err.Error())
		}
		if profilePictureFile != nil {
			_, err = profilePictureFile.Open()
			if err != nil {
				fmt.Printf("B")
				fmt.Printf(err.Error())
			}
			if err != nil {
				fmt.Printf(err.Error())
			}
			profilePictureUrl, err = utils.UploadFile(profilePictureFile)
			if err != nil {
				fmt.Printf(err.Error())
			}
		}

		dob, err := time.Parse("", dobString)
		if err != nil {
			fmt.Printf(err.Error())
		}
		data := &dtosRequest.UpdateUserProfileRequestFORMDTO{
			FirstName:         firstName,
			LastName:          lastName,
			ProfilePictureUrl: profilePictureUrl,
		}
		if !dob.IsZero() {
			data.DOB = &dob
		}
		if profession != "" {
			data.Profession = &profession
		}
		err = validate.Struct(data)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&RegisterUserErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
			})
		}
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		payload, ok := claims["sub"].(map[string]interface{})
		if !ok {
			return c.Status(401).JSON(&UpdateUserProfileErrorResponse{
				Message: "Failed to complete user authentication",
				Status:  401,
			})
		}
		id := payload["ID"].(string)
		entityUser, err := commands.UpdateUserProfile(&dtosCommand.UpdateUserProfileCommandData{
			ID:                id,
			FirstName:         data.FirstName,
			LastName:          data.LastName,
			Profession:        &profession,
			ProfilePictureUrl: data.ProfilePictureUrl,
		})

		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(&utils.DefaultErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadGateway,
			})
		}
		return c.Status(fiber.StatusOK).JSON(&UpdateUserProfileSuccessResponse{
			Message: "Profile updated successfully",
			Status:  fiber.StatusOK,
			Data: &UpdateUserProfileResponseData{
				ID:                entityUser.ID,
				FirstName:         entityUser.FirstName,
				LastName:          entityUser.LastName,
				ProfilePictureUrl: entityUser.ProfilePictureUrl,
				Profession:        *entityUser.Profession,
			},
		})
	}
	return nil
}

type UpdateUserProfileResponseData struct {
	ID                string    `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	ProfilePictureUrl string    `json:"profile_picture_url"`
	Profession        string    `json:"profession"`
	DOB               time.Time `json:"dob"`
}

type UpdateUserProfileSuccessResponse struct {
	Message string                         `json:"message"`
	Status  int                            `json:"status"`
	Data    *UpdateUserProfileResponseData `json:"data"`
}

package handlers

import (
	"strings"

	commands "github.com/datrine/alumni_business/Commands"
	dtosCommand "github.com/datrine/alumni_business/Dtos/Command"
	dtosRequest "github.com/datrine/alumni_business/Dtos/Request"
	utils "github.com/datrine/alumni_business/Utils"
	"github.com/gofiber/fiber/v2"
)

type CreatePostErrorResponse struct {
	Status  int
	Message string
}

// Register: Create new post
//
//	@Summary      Create  new post
//	@Description  Create  new post
//	@Tags         posts
//	@Accept       mpfd
//	@Param        title formData string true "member number"
//	@Param        type_of_post formData string  true "type of post"
//	@Param        text formData string  true "text"
//	@Param	   	  profile_picture formData file true "profile picture"
//	@Produce      json
//	@Success      200  {object}  CreatePostResponseData
//	@Failure      400  {object}  CreatePostErrorResponse
//	@Failure      404  {object}  CreatePostErrorResponse
//	@Failure      500  {object}  CreatePostErrorResponse
//	@Router /posts [post]
func CreatePost(c *fiber.Ctx) error {
	header := &c.Request().Header
	if strings.Contains(strings.ToLower(string(header.ContentType())), "multipart/form-data") {
		title := c.FormValue("title")
		text := c.FormValue("text")
		typeOfPost := c.FormValue("type_of_post")
		mediaMap := make(map[string][]string)
		media, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(&CreatePostErrorResponse{
				Status:  fiber.StatusBadGateway,
				Message: err.Error(),
			})
		}
		for filesKey, files := range media.File {
			for _, file := range files {
				fileUploadURL, err := utils.UploadFile(file)
				if err != nil {
					continue
				}

				mySlice, _ := mediaMap[filesKey]
				mySlice = append(mySlice, fileUploadURL)
				mediaMap[filesKey] = mySlice
			}
		}
		payload, err := utils.GetAuthPayload(c)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(&CreatePostErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadGateway,
			})
		}
		data := &dtosRequest.CreatePost{
			Title:    title,
			Text:     text,
			Type:     typeOfPost,
			Media:    []dtosRequest.MediaDTO{},
			AuthorId: payload.ID,
		}
		err = validate.Struct(data)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(&CreatePostErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadGateway,
			})
		}
		mediaDTO := []dtosCommand.Media{}
		for _, item := range data.Media {
			mediaDTO = append(mediaDTO, dtosCommand.Media{
				URL:  item.URL,
				Type: item.Type,
				Name: item.Name,
			})
		}
		postVTO, err := commands.WritePost(&dtosCommand.WritePost{
			AuthorId: data.AuthorId,
			Title:    data.Title,
			Text:     data.Text,
			Type:     data.Type,
			Media:    mediaDTO,
		})
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(&CreatePostErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
			})
		}
		return c.Status(fiber.StatusCreated).JSON(&CreatePostSuccessResponse{
			Message: "Post created successfully",
			Status:  fiber.StatusCreated,
			Data: &CreatePostResponseData{
				PostID:       postVTO.ID,
				PostText:     postVTO.Text,
				PostTitle:    postVTO.Title,
				PostAuthorId: postVTO.AuthorId,
				PostMedia:    "",
			},
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(&CreatePostErrorResponse{
		Status:  fiber.StatusBadRequest,
		Message: "missing 'multipart/form-data' as content-type header value",
	})
}

type CreatePostResponseData struct {
	PostID       string `json:"post_id"`
	PostText     string `json:"post_text"`
	PostTitle    string `json:"post_title"`
	PostMedia    string `json:"post_media"`
	PostType     string `json:"post_type"`
	PostAuthorId string `json:"post_author_id"`
}

type CreatePostSuccessResponse struct {
	Message string                  `json:"message"`
	Status  int                     `json:"status"`
	Data    *CreatePostResponseData `json:"data"`
}

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

type EditPostErrorResponse struct {
	Status  int
	Message string
}

// GetPosts: Fetch posts
//
//	@Summary      fetch post
//	@Description  fetch post
//	@Tags         posts
//	@Param        sort query string false "asc"
//	@Produce      json
//	@Success      200  {object}  EditPostSuccessResponse
//	@Failure      400  {object}  EditPostErrorResponse
//	@Failure      404  {object}  EditPostErrorResponse
//	@Failure      500  {object}  EditPostErrorResponse
//	@Router /posts [get]
func GetPosts(c *fiber.Ctx) error {
	sort := c.Query("sort")
	postsVTO, err := commands.GetPosts(&dtosCommand.GetPostsDTO{
		Filters: map[string]string{
			"sort": sort,
		},
	})
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(&CreatePostErrorResponse{
			Message: err.Error(),
			Status:  fiber.StatusBadRequest,
		})
	}
	arr := []*FetchPostsResponseData{}
	for _, postModel := range *postsVTO {
		arr = append(arr, &FetchPostsResponseData{
			PostID:       postModel.ID,
			PostTitle:    postModel.Title,
			PostMedia:    "",
			PostText:     postModel.Text,
			PostAuthorId: postModel.AuthorId,
		})
	}
	return c.Status(fiber.StatusOK).JSON(&FetchPostsSuccessResponse{
		Message: "Posts fetched successfully",
		Status:  fiber.StatusOK,
		Data:    &arr,
	})

}

// CreatePost: Create new post
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
//	@Success      200  {object}  CreatePostSuccessResponse
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
				PostType:     postVTO.ContentType,
				PostMedia:    "",
			},
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(&CreatePostErrorResponse{
		Status:  fiber.StatusBadRequest,
		Message: "missing 'multipart/form-data' as content-type header value",
	})
}

// EditPost: Edit post
//
//	@Summary      Edit post
//	@Description  Edit post
//	@Tags         posts
//	@Accept       mpfd
//	@Param        title formData string false "title"
//	@Param        type_of_post formData string  false "type of post"
//	@Param        text formData string  false "text"
//	@Param	   	  profile_picture formData []file false "profile picture"
//	@Produce      json
//	@Success      200  {object}  EditPostSuccessResponse
//	@Failure      400  {object}  EditPostErrorResponse
//	@Failure      404  {object}  EditPostErrorResponse
//	@Failure      500  {object}  EditPostErrorResponse
//	@Router /posts/:id [put]
func EditPost(c *fiber.Ctx) error {
	header := &c.Request().Header
	if strings.Contains(strings.ToLower(string(header.ContentType())), "multipart/form-data") {
		postId := c.Params("id")
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
		data := &dtosRequest.EditPost{
			PostID:   postId,
			Title:    title,
			Text:     text,
			Type:     typeOfPost,
			Media:    []dtosRequest.MediaDTO{},
			AuthorId: payload.ID,
		}
		/*err = validate.Struct(data)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(&EditPostErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadGateway,
			})
		}*/
		mediaDTO := []dtosCommand.Media{}
		for _, item := range data.Media {
			mediaDTO = append(mediaDTO, dtosCommand.Media{
				URL:  item.URL,
				Type: item.Type,
				Name: item.Name,
			})
		}
		postVTO, err := commands.EditPost(&dtosCommand.EditPostDTO{
			AuthorId: data.AuthorId,
			Title:    data.Title,
			Text:     data.Text,
			Type:     data.Type,
			PostID:   data.PostID,
			Media:    mediaDTO,
		})
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(&CreatePostErrorResponse{
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
			})
		}
		return c.Status(fiber.StatusOK).JSON(&EditPostSuccessResponse{
			Message: "Post edited successfully",
			Status:  fiber.StatusOK,
			Data: &EditPostResponseData{
				PostID:       postVTO.ID,
				PostText:     postVTO.Text,
				PostTitle:    postVTO.Title,
				PostAuthorId: postVTO.AuthorId,
				PostType:     postVTO.ContentType,
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

type EditPostResponseData struct {
	PostID       string `json:"post_id"`
	PostText     string `json:"post_text"`
	PostTitle    string `json:"post_title"`
	PostMedia    string `json:"post_media"`
	PostType     string `json:"post_type"`
	PostAuthorId string `json:"post_author_id"`
}

type EditPostSuccessResponse struct {
	Message string                `json:"message"`
	Status  int                   `json:"status"`
	Data    *EditPostResponseData `json:"data"`
}

type FetchPostsSuccessResponse struct {
	Message string                     `json:"message"`
	Status  int                        `json:"status"`
	Data    *[]*FetchPostsResponseData `json:"data"`
}

type FetchPostsResponseData struct {
	PostID       string `json:"post_id"`
	PostText     string `json:"post_text"`
	PostTitle    string `json:"post_title"`
	PostMedia    string `json:"post_media"`
	PostType     string `json:"post_type"`
	PostAuthorId string `json:"post_author_id"`
}

package commands

import (
	dtos "github.com/datrine/alumni_business/Dtos/Command"
	models "github.com/datrine/alumni_business/Models"
	providers "github.com/datrine/alumni_business/Providers"
	"github.com/jaevor/go-nanoid"
)

func GetPosts(data *dtos.GetPostsDTO) (*[]models.Post, error) {
	postModels := &[]models.Post{}

	result := providers.DB.Find(postModels)
	if result.Error != nil {
		return nil, result.Error
	}
	return postModels, nil
}

func WritePost(data *dtos.WritePost) (*models.Post, error) {
	mediaModels := []models.Media{}
	for _, media := range data.Media {
		mediaModels = append(mediaModels, models.Media{
			URL:  media.URL,
			Type: media.Type,
		})
	}
	passwordGenerator, err := nanoid.CustomASCII("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789#$&*@", 10)
	if err != nil {
		return nil, err
	}
	postId := passwordGenerator()
	postModel := &models.Post{
		ID:       postId,
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Media:    mediaModels,
		Text:     data.Text,
	}

	result := providers.DB.Create(postModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return postModel, nil
}

func EditPost(data *dtos.EditPostDTO) (*models.Post, error) {
	mediaModels := []models.Media{}
	for _, media := range data.Media {
		mediaModels = append(mediaModels, models.Media{
			URL:  media.URL,
			Type: media.Type,
		})
	}
	findPostModel := &models.Post{
		AuthorId: data.AuthorId,
		ID:       data.PostId,
	}
	postModel := &models.Post{
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Media:    mediaModels,
		Text:     data.Text,
	}

	result := providers.DB.Model(findPostModel).Updates(postModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return postModel, nil
}

func DeletePost(data *dtos.EditPostDTO) (*models.Post, error) {
	findPostModel := &models.Post{
		AuthorId: data.AuthorId,
		ID:       data.PostId,
	}

	result := providers.DB.Model(findPostModel).First(findPostModel).Delete(findPostModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return findPostModel, nil
}

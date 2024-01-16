package commands

import (
	dtos "github.com/datrine/alumni_business/Dtos/Command"
	models "github.com/datrine/alumni_business/Models"
	providers "github.com/datrine/alumni_business/Providers"
)

func WritePost(data *dtos.WritePost) (*models.Post, error) {
	mediaModels := []models.Media{}
	for _, media := range data.Media {
		mediaModels = append(mediaModels, models.Media{
			URL:  media.URL,
			Type: media.Type,
		})
	}
	postModel := &models.Post{
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

func EditPost(data *dtos.EditPost) (*models.Post, error) {
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

func DeletePost(data *dtos.EditPost) (*models.Post, error) {
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

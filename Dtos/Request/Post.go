package dtos

type CreatePost struct {
	AuthorId string `validate:"required"`
	Title    string `validate:"required"`
	Type     string `validate:"required"`
	Text     string
	Media    []MediaDTO
}

type MediaDTO struct {
	Name string
	URL  string
	Type string
}

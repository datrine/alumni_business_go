package dtos

type EditPost struct {
	PostID   string `validate:"required"`
	AuthorId string `validate:"required"`
	Title    string `validate:"required"`
	Type     string `validate:"required"`
	Text     string
	Media    []MediaDTO
}

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

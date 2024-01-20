package dtos

type GetPostsDTO struct {
	ID      string
	Filters map[string]string
}

type WritePost struct {
	ID       string
	AuthorId string
	Title    string
	Type     string
	Text     string
	Media    []Media
}

type EditPostDTO struct {
	AuthorId string
	PostID   string
	Title    string
	Text     string
	Type     string
	Media    []Media
}

type DeletePostDTO struct {
	AuthorId string
	PostID   string
}
type Media struct {
	URL  string
	Type string
	Name string
}

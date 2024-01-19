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
	ID       string
	AuthorId string
	PostId   string
	Title    string
	Text     string
	Media    []Media
}

type Media struct {
	URL  string
	Type string
	Name string
}

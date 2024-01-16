package dtos

type WritePost struct {
	ID       string
	AuthorId string
	Title    string
	Text     string
	Media    []Media
}

type EditPost struct {
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
}

package model

type BodyComponent struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Images  []string `json:"images"`
}

type Blog struct {
	Id         string        `json:"id"`
	Title      string        `json:"title"`
	Body       BodyComponent `json:"body"`
	Author     string        `json:"author"`
	Timestamp  string        `json:"timestamp"`
	Categories []string      `json:"categories"`
	Tags       []string      `json:"tags"`
}

package models

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

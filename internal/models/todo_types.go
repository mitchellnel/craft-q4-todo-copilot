package models

type TodoItem struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsComplete bool   `json:"is_complete"`
}

type CreateTodoItemRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

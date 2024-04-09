package todo

import (
	"github.com/gin-gonic/gin"
	"mclinton.tech/craft-todo/internal/models"
)

var posts = []models.TodoItem{
	{
		Title:      "Finish this task",
		Content:    "This is a task that needs to be completed",
		IsComplete: false,
	},

	{
		Title:      "Another task",
		Content:    "This is another task that needs to be completed",
		IsComplete: true,
	},

	{
		Title:      "Yet another task",
		Content:    "This is yet another task that needs to be completed",
		IsComplete: false,
	},
}

func GetTodoItemHandler(c *gin.Context) {
	postTitle := c.Param("title")

	for _, post := range posts {
		if post.Title == postTitle {
			c.JSON(200, gin.H{
				"post": post,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error": "Post not found",
	})
}

func GetAllTodoItemshandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func CreateTodoItemHandler(c *gin.Context, request models.CreateTodoItemRequest) {}

func CompleteTodoItemHandler(c *gin.Context) {}

func DeleteTodoItemHandler(c *gin.Context) {}

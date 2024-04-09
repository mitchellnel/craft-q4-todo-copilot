package main

import (
	"github.com/gin-gonic/gin"
	"mclinton.tech/craft-todo/internal/api/todo"
)

func main() {
	server := gin.New()

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	server.GET("/items", todo.GetAllTodoItemshandler)

	// TODO: Register GET /items/:title
	// TODO: Register POST /items
	// TODO: Register PUT /items/:title
	// TOOD: Register DELETE /items/:title

	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}

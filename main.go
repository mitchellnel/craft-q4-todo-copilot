package main

import (
	"github.com/gin-gonic/gin"
	"mclinton.tech/craft-todo/internal/api/posts"
)

func main() {
	server := gin.New()

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	server.GET("/posts", posts.GetAllPostsHandler)

	// TODO: Register GET /posts/:title
	// TOOD: Register POST /posts
	// TOOD: Register DELETE /posts/:title
}

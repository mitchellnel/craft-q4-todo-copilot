package posts

import (
	"github.com/gin-gonic/gin"
	"mclinton.tech/craft-todo/internal/models"
)

var posts = []models.Post{
	models.Post{
		Title:   "First Post",
		Content: "This is the first post",
	},

	models.Post{
		Title:   "Second Post",
		Content: "This is the second post",
	},
}

func GetPostHandler(c *gin.Context) {
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

func GetAllPostsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func CreatePostHandler(c *gin.Context, request models.CreatePostRequest) {}

func DeletePostHandler(c *gin.Context, postTitle string) {}

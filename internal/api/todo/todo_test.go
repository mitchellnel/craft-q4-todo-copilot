package todo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mclinton.tech/craft-todo/internal/api/todo"
)

func TestGetAllPostsHandler(t *testing.T) {}

func TestGetPostHandler(t *testing.T) {
	testCases := []struct {
		name             string
		postTitle        string
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Post found",
			postTitle:        "Finish this task",
			expectedStatus:   200,
			expectedResponse: `{"post":{"title":"Finish this task","content":"This is a task that needs to be completed","is_complete":false}}`,
		},

		{
			name:             "Post not found",
			postTitle:        "Non-existent post",
			expectedStatus:   404,
			expectedResponse: `{"error":"Post not found"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create a mock HTTP request
			req, _ := http.NewRequest("GET", "/posts/"+tc.postTitle, nil)

			// create a new HTTP recorder
			rec := httptest.NewRecorder()

			// create a new server
			server := gin.Default()

			// initialise the API route
			server.GET("/posts/:title", todo.GetTodoItemHandler)

			// make the request, and record the response
			server.ServeHTTP(rec, req)

			// perform assertions
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
		})
	}
}

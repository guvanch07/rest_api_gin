package posts

import (
	"net/http"

	"example.com/rest_api/internal/data"

	"github.com/gin-gonic/gin"
)

// post request
func AddTodo(context *gin.Context) {
	var newTodo data.Todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	// todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

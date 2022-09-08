package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json: "title"`
	Completed bool   `json: "completed"`
}

var todos = []todo{
	{ID: "1", Item: "clean room", Completed: false},
	{ID: "2", Item: "read book", Completed: true},
	{ID: "3", Item: "cook dinner", Completed: false},
}

// get requst
func getTodos(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, todos)
}

// post request
func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:8080")
}

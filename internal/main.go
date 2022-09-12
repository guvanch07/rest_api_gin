package main

import (
	posts "example.com/rest_api/internal/post"

	fetchData "example.com/rest_api/internal/get"
	"example.com/rest_api/internal/update"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", fetchData.GetTodos)
	router.GET("/todos/:id", fetchData.GetTodo)
	router.PATCH("/todos/:id", update.UpdateToggle)
	router.POST("/todos", posts.AddTodo)
	router.Run("localhost:8080")
}

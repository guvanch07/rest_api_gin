package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func updateToggle(context *gin.Context) {
	id := context.Param("id")

	todo, err := findById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

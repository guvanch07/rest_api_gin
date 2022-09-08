package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := findById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}

func findById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("not found")
}

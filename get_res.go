package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get requst
func getTodos(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, todos)
}

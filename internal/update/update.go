package update

import (
	"net/http"

	fetchData "example.com/rest_api/internal/get"
	"github.com/gin-gonic/gin"
)

func UpdateToggle(context *gin.Context) {
	id := context.Param("id")

	todo, err := fetchData.FindById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

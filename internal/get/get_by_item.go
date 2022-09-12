package fetchData

import (
	"errors"
	"net/http"

	"example.com/rest_api/internal/data"
	"github.com/gin-gonic/gin"
)

func GetTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := FindById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}

func FindById(id string) (*data.Todo, error) {
	for i, t := range data.Todos {
		if t.ID == id {
			return &data.Todos[i], nil
		}
	}
	return nil, errors.New("not found")
}

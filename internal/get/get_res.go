package fetchData

import (
	"net/http"

	"example.com/rest_api/internal/data"
	"github.com/gin-gonic/gin"
)

// get requst
func GetTodos(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, data.Todos)
}

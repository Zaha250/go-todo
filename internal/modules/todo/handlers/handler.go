package todo

import (
	"github.com/gin-gonic/gin"
	"go-todo/internal/modules/todo/services"
	"net/http"
)

type TodoHandler struct {
	Service *services.TodoService
}

func NewTodoHandler(s *services.TodoService) *TodoHandler {
	return &TodoHandler{Service: s}
}

func (h *TodoHandler) GetTodoList(c *gin.Context) {
	todos, err := h.Service.GetTodosList()

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": &todos,
	})
}

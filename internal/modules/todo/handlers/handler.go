package todo

import (
	"github.com/gin-gonic/gin"
	todoModels "go-todo/internal/modules/todo/models"
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
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": &todos,
	})
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var data todoModels.CreateTodo

	c.BindJSON(&data)

	newTodoId, createError := h.Service.CreateTodo(data)
	if createError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": createError.Error(),
			"data":  data,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data":   newTodoId,
	})
}

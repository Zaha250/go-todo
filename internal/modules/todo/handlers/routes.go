package todo

import (
	"github.com/gin-gonic/gin"
	"go-todo/internal/modules/todo/repositories"
	"go-todo/internal/modules/todo/services"
)

func RegisterHandlers(router *gin.RouterGroup) {
	h := NewTodoHandler(
		services.NewTodoService(
			repositories.NewTodoRepository("internal/storage/todos.json"),
		),
	)

	routes := router.Group("/todos")

	{
		routes.GET("/", h.GetTodoList)
		routes.GET("/:id", h.GetTodoById)
		routes.POST("/", h.CreateTodo)
		routes.DELETE("/:id", h.DeleteTodo)
		routes.PATCH("/", h.UpdateTodo)
	}
}

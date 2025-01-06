package services

import (
	"errors"
	todoModels "go-todo/internal/modules/todo/models"
	"go-todo/internal/modules/todo/repositories"
)

type TodoService struct {
	Repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{Repo: repo}
}

func (s *TodoService) GetTodosList() ([]todoModels.Todo, error) {
	return s.Repo.GetList()
}

func (s *TodoService) GetTodoById(todoId string) (*todoModels.Todo, error) {
	return s.Repo.GetById(todoId)
}

func (s TodoService) CreateTodo(data todoModels.CreateTodo) error {
	if data.Title == "" {
		return errors.New("введите название задачи")
	}
	return s.Repo.Create(data)
}

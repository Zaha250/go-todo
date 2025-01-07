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

func (s *TodoService) GetTodoById(todoId todoModels.TodoId) (*todoModels.Todo, error) {
	return s.Repo.GetById(todoId)
}

func (s *TodoService) CreateTodo(data todoModels.CreateTodo) (todoModels.TodoId, error) {
	if data.Title == "" {
		return "", errors.New("введите название задачи")
	}
	return s.Repo.Create(data)
}

func (s *TodoService) DeleteTodo(todoId todoModels.TodoId) error {
	return s.Repo.Delete(todoId)
}

func (s *TodoService) UpdateTodo(data todoModels.UpdateTodo) error {
	return s.Repo.Update(data)
}

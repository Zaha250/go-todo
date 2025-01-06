package repositories

import (
	"encoding/json"
	"errors"
	todoModels "go-todo/internal/modules/todo/models"
	"os"
	"time"
)

type TodoRepository struct {
	filePath string
}

func NewTodoRepository(filePath string) *TodoRepository {
	return &TodoRepository{filePath}
}

func (r *TodoRepository) GetList() ([]todoModels.Todo, error) {
	file, err := os.Open(r.filePath)

	if err != nil {
		return nil, err
	}
	/* если файл пустой, то возвращаем пустой массив */
	fileInfo, _ := file.Stat()

	if fileInfo.Size() == 0 {
		return []todoModels.Todo{}, nil
	}

	defer file.Close()

	var todos []todoModels.Todo

	errorDecode := json.NewDecoder(file).Decode(&todos)
	if errorDecode != nil {
		return nil, errorDecode
	}

	return todos, nil
}

func (r *TodoRepository) GetById(id string) (*todoModels.Todo, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var todos []todoModels.Todo

	errorDecode := json.NewDecoder(file).Decode(&todos)
	if errorDecode != nil {
		return nil, errorDecode
	}

	for _, todoItem := range todos {
		if todoItem.Id == id {
			return &todoItem, nil
		}
	}
	return nil, errors.New("todo not found")
}

func (r *TodoRepository) writeTodos(newTodos []todoModels.Todo) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(newTodos)
}

func (r *TodoRepository) Create(data todoModels.CreateTodo) error {
	todos, err := r.GetList()
	if err != nil {
		return err
	}

	newTodo := todoModels.Todo{
		Id:        time.Now().UTC().String(),
		Title:     data.Title,
		Completed: false,
	}

	todos = append(todos, newTodo)
	return r.writeTodos(todos)
}

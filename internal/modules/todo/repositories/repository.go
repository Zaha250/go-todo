package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	todoModels "go-todo/internal/modules/todo/models"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
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

func (r *TodoRepository) GetById(id todoModels.TodoId) (*todoModels.Todo, error) {
	todos, err := r.GetList()
	if err != nil {
		return nil, err
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

func (r *TodoRepository) dumpDB() error {
	sourceFile, err := os.Open(r.filePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	sourceFileName := strings.TrimSuffix(sourceFile.Name(), path.Ext(r.filePath))
	destinationFileName := sourceFileName + "_" + time.Now().UTC().Format(time.RFC3339) + ".json"

	destinationFile, err := os.Create(destinationFileName)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) Create(data todoModels.CreateTodo) (todoModels.TodoId, error) {
	todos, err := r.GetList()
	if err != nil {
		return "", err
	}

	/* Делаем dump перед каждым добавлением записи */
	dumpErr := r.dumpDB()
	if dumpErr != nil {
		return "", dumpErr
	}

	newTodo := todoModels.Todo{
		Id:        todoModels.TodoId(strconv.FormatInt(time.Now().UnixNano(), 10)),
		Title:     data.Title,
		Completed: false,
	}

	todos = append(todos, newTodo)
	writesErr := r.writeTodos(todos)

	return newTodo.Id, writesErr
}

func (r *TodoRepository) Delete(id todoModels.TodoId) error {
	allTodos, err := r.GetList()
	if err != nil {
		return err
	}

	/* Делаем dump перед каждым удалением записи */
	dumpErr := r.dumpDB()
	if dumpErr != nil {
		return dumpErr
	}

	var filteredTodos []todoModels.Todo
	isFound := false

	for _, todo := range allTodos {
		if todo.Id != id {
			filteredTodos = append(filteredTodos, todo)
		} else {
			isFound = true
		}
	}
	if !isFound {
		return errors.New(string("Не найдена задача с таким id " + id))
	}
	return r.writeTodos(filteredTodos)
}

func (r *TodoRepository) Update(data todoModels.UpdateTodo) error {
	todos, err := r.GetList()
	if err != nil {
		return err
	}

	dumpErr := r.dumpDB()
	if dumpErr != nil {
		return dumpErr
	}

	for i, todo := range todos {
		if todo.Id == data.Id {
			newTodo := todoModels.Todo{
				Id:        todo.Id,
				Title:     todo.Title,
				Completed: data.Completed,
			}
			todos[i] = newTodo

			err := r.writeTodos(todos)
			if err != nil {
				return err
			}

			return nil
		}
	}
	return fmt.Errorf("задача с номером %s не найдена", data.Id)
}

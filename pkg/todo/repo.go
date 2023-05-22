package todo

import (
	"encoding/json"
	"io"
	"os"
	"todolist/pkg/models"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (t *TodoRepository) GetTodoList() (todoList models.TodoList, err error) {

	todoFile, err := os.Open("database/todo.json")
	if err != nil {
		return todoList, err
	}
	defer todoFile.Close()

	byteValue, err := io.ReadAll(todoFile)
	if err != nil {
		return todoList, err
	}

	if err = json.Unmarshal(byteValue, &todoList); err != nil {
		return todoList, err
	}

	return todoList, nil
}

func (t *TodoRepository) WriteTodoList(todoList models.TodoList) error {

	todoFile, err := os.OpenFile("database/todo.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer todoFile.Close()

	jsonData, err := json.Marshal(todoList)
	if err != nil {
		return err
	}

	if _, err := todoFile.Write(jsonData); err != nil {
		return err
	}

	return nil
}

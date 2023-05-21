package todo

import (
	"errors"
	"log"
	"strings"
	"todolist/pkg/domains"
	"todolist/pkg/models"
	"todolist/utils"
)

type TodoUsecase struct {
	repo domains.TodoRepository
}

func NewTodoUsecase(r domains.TodoRepository) domains.TodoUsecase {
	return &TodoUsecase{
		repo:  r,
	}
}

func (u *TodoUsecase) GetTodo(sortBy string) (todoList models.TodoList, err error) {
	todoList, err = u.repo.GetTodoList()
	if err != nil {
		return
	}

	switch strings.ToLower(sortBy) {
	case "title":
		todoList.SortByTitle()
	case "date":
		todoList.SortByDate()
	case "status":
		todoList.SortByStatus()
	}

	return
}

func (u *TodoUsecase) SearchTodo(title string, desc string) (todoList models.TodoList, err error) {
	todoList, err = u.repo.GetTodoList()
	if err != nil {
		return
	}

	if title != "" || desc != "" {
		var SearchTodoList models.TodoList
		title = strings.ToLower(title)
		desc = strings.ToLower(desc)
		for _, todo := range todoList.TodoList {
			lowerTitle := strings.ToLower(todo.Title)
			lowerDesc := strings.ToLower(utils.GetStringValue(todo.Description))
			if strings.Contains(lowerTitle, title) || (lowerDesc != "" && strings.Contains(lowerDesc, desc)) {
				SearchTodoList.TodoList = append(SearchTodoList.TodoList, todo)
			}
		}
		return SearchTodoList, nil
	}
	return
}

func (u *TodoUsecase) CreateTodo(todoReq models.TodoDetail) error {
	todoList, err := u.repo.GetTodoList()
	if err != nil {
		return err
	}

	for _, todo := range todoList.TodoList {
		if todo.ID == todoReq.ID {
			return errors.New("same Id")
		}
	}

	log.Println("todoReq : ", todoReq)
	todoList.TodoList = append(todoList.TodoList, todoReq)
	if err := u.repo.WriteTodoList(todoList); err != nil {
		return err
	}

	return nil
}

func (u *TodoUsecase) UpdateTodo(updateData models.TodoDetail) error {
	todoList, err := u.repo.GetTodoList()
	if err != nil {
		return err
	}

	todoId := updateData.ID
	for index, todoData := range todoList.TodoList {
		if todoData.ID == todoId {
			todoList.TodoList[index] = updateData
			if err := u.repo.WriteTodoList(todoList); err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("id not found")
}

package domains

import "todolist/pkg/models"

type TodoUsecase interface {
	GetTodo(sortBy string) (todoList models.TodoList, err error)
	SearchTodo(title string, desc string) (todoList models.TodoList, err error)
	CreateTodo(todoReq models.TodoDetail) error
	UpdateTodo(updateData models.TodoDetail) error
}

type TodoRepository interface {
	GetTodoList() (todoList models.TodoList, err error)
	WriteTodoList(todoList models.TodoList) error
}

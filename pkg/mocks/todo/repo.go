package mocks

import (
	"todolist/pkg/models"

	mock "github.com/stretchr/testify/mock"
)

type TodoRepository struct {
	mock.Mock
}

func (_m *TodoRepository) GetTodoList() (models.TodoList, error) {
	ret := _m.Called()

	var r0 models.TodoList
	if rf, ok := ret.Get(0).(func() models.TodoList); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.TodoList)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TodoRepository) WriteTodoList(todoList models.TodoList) error {
	ret := _m.Called(todoList)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.TodoList) error); ok {
		r0 = rf(todoList)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
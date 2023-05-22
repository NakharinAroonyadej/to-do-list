package mocks

import (
	"todolist/pkg/models"

	"github.com/stretchr/testify/mock"
)

type TodoUsecase struct {
	mock.Mock
}

func (_m * TodoUsecase) GetTodo(sortBy string) (todoList models.TodoList, err error) {
	ret := _m.Called(sortBy)

	var r0 models.TodoList
	if rf, ok := ret.Get(0).(func(string) models.TodoList); ok {
		r0 = rf(sortBy)
	} else {
		r0 = ret.Get(0).(models.TodoList)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sortBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m * TodoUsecase) SearchTodo(title string, desc string) (todoList models.TodoList, err error) {
	ret := _m.Called(title, desc)

	var r0 models.TodoList
	if rf, ok := ret.Get(0).(func(string, string) models.TodoList); ok {
		r0 = rf(title, desc)
	} else {
		r0 = ret.Get(0).(models.TodoList)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(title, desc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m * TodoUsecase) CreateTodo(todoReq models.TodoDetail) error {
	ret := _m.Called(todoReq)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.TodoDetail) error); ok {
		r0 = rf(todoReq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m * TodoUsecase) UpdateTodo(updateData models.TodoDetail) error {
	ret := _m.Called(updateData)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.TodoDetail) error); ok {
		r0 = rf(updateData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

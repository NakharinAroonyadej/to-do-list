package todo_test

import (
	"errors"
	"testing"
	"time"
	mocks "todolist/pkg/mocks/todo"
	"todolist/pkg/models"
	"todolist/pkg/todo"
	"todolist/utils"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodoUsecase := new(mocks.TodoUsecase)
	mockTodoList := models.TodoList{
		TodoList: []models.TodoDetail{
			{
				ID:          uuid.New(),
				Title:       "Test GetTodo",
				Description: nil,
				CreatedAt:   time.Now(),
				Image:       nil,
				Status:      "IN_PROGRESS",
			},
		},
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(mockTodoList, nil).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		got, err := usecase.GetTodo("")

		assert.NoError(t, err)
		assert.Equal(t, mockTodoList, got)

		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("failed-error", func(t *testing.T) {
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(models.TodoList{}, errors.New("repository error")).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		got, err := usecase.GetTodo("")

		assert.Error(t, err)
		assert.Equal(t, models.TodoList{}, got)

		mockTodoUsecase.AssertExpectations(t)
	})
}

func TestSearchTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodoUsecase := new(mocks.TodoUsecase)
	mockSearchByTitle := models.TodoDetail{
		ID:          uuid.New(),
		Title:       "Search By Title",
		Description: utils.NewString("Test SearchTodo"),
		CreatedAt:   time.Now(),
		Image:       nil,
		Status:      "IN_PROGRESS",
	}
	mockSearchByDesc := models.TodoDetail{
		ID:          uuid.New(),
		Title:       "Test SearchTodo",
		Description: utils.NewString("Search By Description"),
		CreatedAt:   time.Now(),
		Image:       nil,
		Status:      "IN_PROGRESS",
	}
	mockTodoList := models.TodoList{
		TodoList: []models.TodoDetail{
			{
				ID:          uuid.New(),
				Title:       "Test SearchTodo 2",
				Description: nil,
				CreatedAt:   time.Now(),
				Image:       nil,
				Status:      "IN_PROGRESS",
			},
			mockSearchByTitle,
			mockSearchByDesc,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(mockTodoList, nil).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		got, err := usecase.SearchTodo("", "")

		assert.NoError(t, err)
		assert.Equal(t, mockTodoList, got)

		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("search by title", func(t *testing.T) {
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(mockTodoList, nil).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		got, err := usecase.SearchTodo("title", "")

		assert.NoError(t, err)
		assert.Equal(t, models.TodoList{
			TodoList: []models.TodoDetail{mockSearchByTitle},
		}, got)

		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("search by description", func(t *testing.T) {
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(mockTodoList, nil).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		got, err := usecase.SearchTodo("", "description")

		assert.NoError(t, err)
		assert.Equal(t, models.TodoList{
			TodoList: []models.TodoDetail{mockSearchByDesc},
		}, got)
		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("failed-error", func(t *testing.T) {
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(models.TodoList{}, errors.New("repository error")).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		got, err := usecase.SearchTodo("", "")

		assert.Error(t, err)
		assert.Equal(t, models.TodoList{}, got)

		mockTodoUsecase.AssertExpectations(t)
	})
}

func TestCreateTodo(t *testing.T) {
	mockTodoUsecase := new(mocks.TodoUsecase)
	var fakeTodo models.TodoDetail
	gofakeit.Struct(fakeTodo)
	mockTodoList := models.TodoList{
		TodoList: []models.TodoDetail{
			fakeTodo,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(models.TodoList{}, nil).Once()
		mockTodoRepo.On("WriteTodoList", mock.AnythingOfType("TodoList")).Return(nil).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		err := usecase.CreateTodo(fakeTodo)

		assert.NoError(t, err)

		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("same id", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(mockTodoList, nil).Once()
		mockTodoRepo.On("WriteTodoList", mock.Anything).Return(nil).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		err := usecase.CreateTodo(fakeTodo)

		assert.Error(t, err)

		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("repo write error", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(models.TodoList{}, nil).Once()
		mockTodoRepo.On("WriteTodoList", mock.Anything).Return(errors.New("Write error")).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		err := usecase.CreateTodo(fakeTodo)

		assert.Error(t, err)

		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("repo read error", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(models.TodoList{}, errors.New("read repo error")).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		err := usecase.CreateTodo(fakeTodo)

		assert.Error(t, err)

		mockTodoUsecase.AssertExpectations(t)
	})

}

func TestUpdateTodo(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodoUsecase := new(mocks.TodoUsecase)
	var fakeTodo models.TodoDetail
	gofakeit.Struct(fakeTodo)
	mockTodoList := models.TodoList{
		TodoList: []models.TodoDetail{
			fakeTodo,
		},
	}


	t.Run("success", func(t *testing.T) {
		mockUpdate := models.TodoDetail{
			ID:          fakeTodo.ID,
			Title:       "New Title",
			Description: utils.NewString("New Description"),
			CreatedAt:   time.Now(),
			Status:      "COMPLETE",
		}

		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(mockTodoList, nil).Once()
		mockTodoRepo.On("WriteTodoList", mock.AnythingOfType("TodoList")).Return(nil).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		err := usecase.UpdateTodo(mockUpdate)

		assert.NoError(t, err)

		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("id not found", func(t *testing.T) {
		mockUpdate := models.TodoDetail{
			ID:          uuid.New(),
		}

		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(mockTodoList, nil).Once()
		mockTodoRepo.On("WriteTodoList", mock.AnythingOfType("TodoList")).Return(nil).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		err := usecase.UpdateTodo(mockUpdate)

		assert.Error(t, err)

		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("repo read error", func(t *testing.T) {
		mockUpdate := models.TodoDetail{
			ID:          uuid.New(),
		}

		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(models.TodoList{}, errors.New("read repo error")).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		err := usecase.UpdateTodo(mockUpdate)

		assert.Error(t, err)

		mockTodoUsecase.AssertExpectations(t)
	})

	t.Run("repo write error", func(t *testing.T) {
		mockUpdate := models.TodoDetail{
			ID:          uuid.New(),
		}
		mockTodoRepo.On("GetTodoList", mock.AnythingOfType("string")).Return(models.TodoList{}, nil).Once()
		mockTodoRepo.On("WriteTodoList", mock.Anything).Return(errors.New("Write error")).Once()

		usecase := todo.NewTodoUsecase(mockTodoRepo)

		err := usecase.UpdateTodo(mockUpdate)

		assert.Error(t, err)

		mockTodoUsecase.AssertExpectations(t)
	})
}

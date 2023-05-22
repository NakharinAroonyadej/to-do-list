package todo_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	mocks "todolist/pkg/mocks/todo"
	"todolist/pkg/models"
	"todolist/pkg/todo"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type todoReq struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Status    string `json:"status"`
}

func TestGetTodoHandler(t *testing.T) {
	mockUsecase := new(mocks.TodoUsecase)
	var fakeTodo models.TodoDetail
	gofakeit.Struct(&fakeTodo)
	mockTodoList := models.TodoList{
		TodoList: []models.TodoDetail{
			fakeTodo,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockUsecase.On("GetTodo", mock.AnythingOfType("string")).Return(mockTodoList, nil).Once()

		app := gin.New()
		gin.SetMode(gin.ReleaseMode)
		handler := todo.NewTodoHandler(mockUsecase)
		app.GET("/get", handler.GetTodo)

		req, err := http.NewRequest("GET", "/get", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

func TestSearchTodoHandler(t *testing.T) {
	mockUsecase := new(mocks.TodoUsecase)
	var fakeTodo models.TodoDetail
	gofakeit.Struct(&fakeTodo)
	mockTodoList := models.TodoList{
		TodoList: []models.TodoDetail{
			fakeTodo,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockUsecase.On("SearchTodo", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockTodoList, nil).Once()

		app := gin.New()
		gin.SetMode(gin.ReleaseMode)
		handler := todo.NewTodoHandler(mockUsecase)
		app.GET("/search", handler.SearchTodo)

		req, err := http.NewRequest("GET", `/search`, nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

func TestCreateTodoHandler(t *testing.T) {
	mockUsecase := new(mocks.TodoUsecase)
	t.Run("success", func(t *testing.T) {
		mockUsecase.On("CreateTodo", mock.Anything).Return(nil).Once()

		gin.SetMode(gin.TestMode)
		handler := todo.NewTodoHandler(mockUsecase)

		todoReq := todoReq{
			ID:        "00000000-0000-0000-0000-000000000003",
			Title:     "Test",
			CreatedAt: "2019-10-12T14:20:50.52+07:00",
			Status:    "COMPLETE",
		}

		jsonValue, _ := json.Marshal(todoReq)

		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonValue))
		c.Request.Header.Set("Content-Type", "application/json")
		r.POST("/create", handler.CreateTodo)
		r.ServeHTTP(w, c.Request)

		var todoResp models.ResponseForm
		json.Unmarshal(w.Body.Bytes(), &todoResp)
		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

func TestUpdateTodoHandler(t *testing.T) {
	mockUsecase := new(mocks.TodoUsecase)
	t.Run("success", func(t *testing.T) {
		mockUsecase.On("UpdateTodo", mock.Anything).Return(nil).Once()

		gin.SetMode(gin.TestMode)
		handler := todo.NewTodoHandler(mockUsecase)

		todoReq := todoReq{
			ID:        "00000000-0000-0000-0000-000000000003",
			Title:     "Test",
			CreatedAt: "2019-10-12T14:20:50.52+07:00",
			Status:    "COMPLETE",
		}

		jsonValue, _ := json.Marshal(todoReq)

		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPatch, "/update", bytes.NewBuffer(jsonValue))
		c.Request.Header.Set("Content-Type", "application/json")
		r.PATCH("/update", handler.UpdateTodo)
		r.ServeHTTP(w, c.Request)

		var todoResp models.ResponseForm
		json.Unmarshal(w.Body.Bytes(), &todoResp)
		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})
}

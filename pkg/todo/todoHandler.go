package todo

import (
	"net/http"
	"todolist/pkg/domains"
	"todolist/pkg/models"
	"todolist/validates"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUsecase domains.TodoUsecase
}

func NewTodoHandler(usecase domains.TodoUsecase) TodoHandler {
	return TodoHandler{
		todoUsecase: usecase,
	}
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	sortBy := c.Param("sortBy")

	if err := validates.ValidateGetTodosRequest(sortBy); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	todoList, err := h.todoUsecase.GetTodo(sortBy)
	if err != nil {
		c.AbortWithError(http.StatusExpectationFailed, err)
	}

	c.JSON(http.StatusOK, todoList)
}

func (h *TodoHandler) SearchTodo(c *gin.Context) {
	title := c.Query("title")
	desc := c.Query("description")

	todoList, err := h.todoUsecase.SearchTodo(title, desc)
	if err != nil {
		c.AbortWithError(http.StatusExpectationFailed, err)
	}

	c.JSON(http.StatusOK, todoList)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todoReq models.TodoDetail
	if err := c.Bind(&todoReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := validates.ValidateTodoRequest(todoReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.todoUsecase.CreateTodo(todoReq); err != nil {
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}

	c.JSON(http.StatusOK, "create complete!")
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var todoReq models.TodoDetail
	if err := c.Bind(&todoReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := validates.ValidateTodoRequest(todoReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.todoUsecase.UpdateTodo(todoReq); err != nil {
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}

	c.JSON(http.StatusOK, "update complete!")
}

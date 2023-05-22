package todo

import (
	"net/http"
	"todolist/pkg/domains"
	"todolist/pkg/models"
	"todolist/utils"
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
		c.JSON(http.StatusBadRequest, utils.NewErrorResponseForm(http.StatusBadRequest, err.Error()))
		return
	}

	todoList, err := h.todoUsecase.GetTodo(sortBy)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, utils.NewErrorResponseForm(http.StatusExpectationFailed, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponseForm(todoList))
}

func (h *TodoHandler) SearchTodo(c *gin.Context) {
	title := c.Query("title")
	desc := c.Query("description")

	todoList, err := h.todoUsecase.SearchTodo(title, desc)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, utils.NewErrorResponseForm(http.StatusExpectationFailed, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponseForm(todoList))
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todoReq models.TodoDetail
	if err := c.Bind(&todoReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponseForm(http.StatusBadRequest, err.Error()))
		return
	}

	if err := validates.ValidateTodoRequest(todoReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponseForm(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.todoUsecase.CreateTodo(todoReq); err != nil {
		c.JSON(http.StatusExpectationFailed, utils.NewErrorResponseForm(http.StatusExpectationFailed, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponseForm(todoReq))
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var todoReq models.TodoDetail
	if err := c.Bind(&todoReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponseForm(http.StatusBadRequest, err.Error()))
		return
	}

	if err := validates.ValidateTodoRequest(todoReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponseForm(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.todoUsecase.UpdateTodo(todoReq); err != nil {
		c.JSON(http.StatusExpectationFailed, utils.NewErrorResponseForm(http.StatusExpectationFailed, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponseForm(todoReq))
}

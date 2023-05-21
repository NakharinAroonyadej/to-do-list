package models

import (
	"sort"
	"time"

	"github.com/google/uuid"
)

type TodoList struct {
	TodoList []TodoDetail `json:"todoList"`
}

type TodoDetail struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	Image       *string   `json:"image"`
	Status      string    `json:"status"`
}

func (list *TodoList) SortByTitle() {
	sort.Slice(list.TodoList, func(i, j int) bool {
		return list.TodoList[i].Title < list.TodoList[j].Title
	})
}

func (list *TodoList) SortByDate() {
	sort.Slice(list.TodoList, func(i, j int) bool {
		return list.TodoList[i].CreatedAt.Unix() < list.TodoList[j].CreatedAt.Unix()
	})
}

func (list *TodoList) SortByStatus() {
	sort.Slice(list.TodoList, func(i, j int) bool {
		return list.TodoList[i].Status < list.TodoList[j].Status
	})
}

package validates

import (
	"encoding/base64"
	"errors"
	"time"
	"todolist/pkg/models"
	"todolist/utils"

	"github.com/google/uuid"
)

func ValidateGetTodosRequest(sortByReq string) error {
	if sortByReq == "" {
		return nil
	}
	sortByList := []string{"Title", "Date", "Status"}
	for _, sortBy := range sortByList {
		if sortBy == sortByReq {
			return nil
		}
	}
	return errors.New("sortby param invalid")
}

func ValidateTodoRequest(todo models.TodoDetail) error {
	if !HasStringValue(todo.ID.String()) || !HasStringValue(todo.Title) || !HasStringValue(todo.CreatedAt.String()) || !HasStringValue(todo.Status) {
		return errors.New("id, title, date, status must be required")
	}
	if _, err := uuid.Parse(todo.ID.String()); err != nil {
		return errors.New("id must be uuid format")
	}
	if len(todo.Title) >= 100 {
		return errors.New("title must not over 100 characters")
	}
	if _, err := time.Parse(time.RFC3339, todo.CreatedAt.Format(time.RFC3339)); err != nil {
		return errors.New("date must be RFC3339 with timezone")
	}
	if (todo.Status != "IN_PROGRESS" && todo.Status != "COMPLETE") {
		return errors.New("status must be IN_PROGRESS or COMPLETE")
	}
	if todo.Image != nil {
		if _, err := base64.StdEncoding.DecodeString(utils.GetStringValue(todo.Image)); err != nil {
			return errors.New("image must be base64 format")
		}
	}
	return nil
}

func HasStringValue(s string) bool {
	return s != ""
}

package utils

import (
	"net/http"
	"todolist/pkg/models"
)

func NewSuccessResponseForm(data interface{}) models.ResponseForm {
	return models.ResponseForm{
		Success: true,
		Data: data,
		ErrorsCode: http.StatusOK,
		Messages: "Success",
	}
}

func NewErrorResponseForm(errorCode int, errorMessage string) models.ResponseForm {
	return models.ResponseForm{
		Success: false,
		Data: nil,
		ErrorsCode: errorCode,
		Messages: errorMessage,
	}
}
package models

type ResponseForm struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Messages   string      `json:"messages,omitempty"`
	ErrorsCode int         `json:"errors,omitempty"`
}

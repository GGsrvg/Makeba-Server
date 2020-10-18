package models

type Response struct {
	Message      *string
	ErrorMessage *string // this only for developers
	Data         interface{}
}

func NewRespose(message *string, errorMessage *string, data interface{}) *Response {
	return &Response{Message: message, ErrorMessage: errorMessage, Data: data}
}

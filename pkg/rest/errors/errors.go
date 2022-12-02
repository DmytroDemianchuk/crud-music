package errors

import (
	"net/http"
)

type NotFoundErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (n NotFoundErr) RuntimeError() {
	//TODO implement me
	panic("implement me")
}

func NewNotFoundErr(message string) NotFoundErr {
	return NotFoundErr{Code: http.StatusNotFound, Message: message}
}

type BadRequestErr struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

func NewBadRequestErr(mgs string, fields map[string]string) BadRequestErr {
	return BadRequestErr{Code: http.StatusBadRequest, Message: mgs, Fields: fields}
}

type InternalServerErr struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func NewInternalServerErr() InternalServerErr {
	return InternalServerErr{Code: http.StatusInternalServerError, Message: "something went wrong"}
}

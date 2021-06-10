package rest_errors

import (
	"errors"
	"fmt"
	"net/http"
)

type RestError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restError struct {
	message string `json:"message"`
	status  int    `json:"status"`
	error   string `json:"error"`
	causes []interface{} `json:"causes"`
}

func (e restError) Error() string {
	return fmt.Sprintf("message :%s - status: %d - error: %s - causes [ %v ]",
		e.message, e.status, e.error, e.causes)
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestError(message string, status int, err string, causes []interface{}) RestError {
	return restError{
		message: message,
		status: status,
		error: err,
		causes: causes,
	}
}

func NewUnauthorizedError(message string) RestError {
	return restError{
		message: message,
		status: http.StatusUnauthorized,
		error: "unauthorized",
	}
}

func NewBadRequestError(message string) RestError {
	return &restError{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

func NewNotFoundError(message string) RestError {
	return &restError{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) RestError {
	result := &restError{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}

func (e restError) Message() string {
	return e.message
}

func (e restError) Status() int {
	return e.status
}

func (e restError) Causes() []interface{} {
	return e.causes
}
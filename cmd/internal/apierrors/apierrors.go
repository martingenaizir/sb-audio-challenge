package apierrors

import (
	"fmt"
	"net/http"
)

type CauseList []any

type ApiError interface {
	Message() string
	Code() int
	Status() string
	Cause() CauseList
	Error() string
}

type apiError struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    int       `json:"code"`
	ErrorStatus  string    `json:"status"`
	ErrorCause   CauseList `json:"cause,omitempty"`
}

func (a apiError) Message() string {
	return a.ErrorMessage
}

func (a apiError) Code() int {
	return a.ErrorCode
}

func (a apiError) Status() string {
	return a.ErrorStatus
}

func (a apiError) Cause() CauseList {
	return a.ErrorCause
}

func (a apiError) Error() string {
	return fmt.Sprintf("Message: %s; Error code: %d; Error status: %s; Cause: %v",
		a.ErrorMessage, a.ErrorCode, a.ErrorStatus, a.ErrorCause)
}

func BadRequestError(message string) ApiError {
	return apiError{
		ErrorMessage: message,
		ErrorCode:    http.StatusBadRequest,
		ErrorStatus:  http.StatusText(http.StatusBadRequest),
	}
}

func NotFound(message string) ApiError {
	return apiError{
		ErrorMessage: message,
		ErrorCode:    http.StatusNotFound,
		ErrorStatus:  http.StatusText(http.StatusNotFound),
	}
}

func ContentTooLargeError(message string) ApiError {
	return apiError{
		ErrorMessage: message,
		ErrorCode:    http.StatusRequestEntityTooLarge,
		ErrorStatus:  http.StatusText(http.StatusRequestEntityTooLarge),
	}
}

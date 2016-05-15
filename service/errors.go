package service

import (
	"fmt"
	"net/http"
)

// HandlerError wraps an error for the handlers, allows custom error codes
type HandlerError struct {
	Err  error
	Code int
}

// Error implements the Error interface
func (he *HandlerError) Error() string {
	return fmt.Sprintf("%s", he.Err)
}

// NewError is a helper function for creating a new HandlerError wrapping an error
func NewError(e error) *HandlerError {
	return &HandlerError{Err: e, Code: http.StatusInternalServerError}
}

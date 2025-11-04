package ginrouter

import (
	"fmt"
)

// BadRequestError represents a bad request error with a message.
type BadRequestError struct {
	Message string `json:"message"`
}

type CustomError interface {
	StatusCode() int
	Error() string
}

// NewIDShouldBeIntError creates a new BadRequestError for invalid ID type.
func NewIDShouldBeIntError(param string) *BadRequestError {
	return &BadRequestError{
		Message: fmt.Sprintf("bad request: %s should be an integer", param),
	}
}

func (e *BadRequestError) StatusCode() int {
	return 400
}

func (e *BadRequestError) Error() string {
	return e.Message
}
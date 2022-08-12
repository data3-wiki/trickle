package model

import "fmt"

// InputValidationError is used for all errors related to input validation.
type InputValidationError struct {
	message string
}

// NewInputValidationError constructs a new input validation error.
func NewInputValidationError(format string, a ...interface{}) *InputValidationError {
	return &InputValidationError{
		message: fmt.Sprintf(format, a...),
	}
}

// Error returns the error message.
func (e *InputValidationError) Error() string {
	return e.message
}

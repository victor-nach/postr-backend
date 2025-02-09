package domain

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
)

const errorStatus = "error"

type DomainError struct {
	Status      string            `json:"status"`
	Code        string            `json:"code"`
	Message     string            `json:"message"`
	FieldErrors map[string]string `json:"fieldErrors,omitempty"`
}

func (e DomainError) Error() string {
	if len(e.FieldErrors) > 0 {
		return fmt.Sprintf("%s: %s (field errors: %v)", e.Code, e.Message, e.FieldErrors)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// WithFieldErrors attaches validation errors to a DomainError
func (e DomainError) WithFieldErrors(errs validation.Errors) DomainError {
	fieldErrors := make(map[string]string, len(errs))
	for field, err := range errs {
		fieldErrors[field] = err.Error()
	}
	e.FieldErrors = fieldErrors
	return e
}

func (e DomainError) Is(target error) bool {
	if t, ok := target.(DomainError); ok {
		return e.Code == t.Code
	}

	if t, ok := target.(*DomainError); ok {
		return e.Code == t.Code
	}

	return false
}

var (
	ErrInternalServer = DomainError{
		Status:  errorStatus,
		Code:    "APP-500",
		Message: "Internal server error - Unable to handle request",
	}

	ErrInvalidInput = DomainError{
		Status:  errorStatus,
		Code:    "APP-400",
		Message: "Invalid input data",
	}

	ErrUserNotFound = DomainError{
		Status:  errorStatus,
		Code:    "USR-404001",
		Message: "User not found",
	}

	ErrPostNotFound = DomainError{
		Status:  errorStatus,
		Code:    "PST-404001",
		Message: "Post not found",
	}

	ErrCreateUser = DomainError{
		Status:  errorStatus,
		Code:    "USR-400101",
		Message: "Failed to create user",
	}
)

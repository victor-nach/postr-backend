package domain

import "github.com/go-ozzo/ozzo-validation/v4"

type DomainError struct {
	Code        string            `json:"code"`
	Message     string            `json:"message"`
	FieldErrors map[string]string `json:"fieldErrors,omitempty"`
}

// Error implements error.
func (e DomainError) Error() string {
	panic("unimplemented")
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

var (
	ErrInternalServer = DomainError{
		Code:    "APP-500",
		Message: "Internal server error - Unable to handle request",
	}

	ErrInvalidInput = DomainError{
		Code:    "APP-400",
		Message: "Invalid input data",
	}

	ErrUserNotFound = DomainError{
		Code:    "USR-404001",
		Message: "User not found",
	}

	ErrPostNotFound = DomainError{
		Code:    "PST-404001",
		Message: "Post not found",
	}

	ErrCreateUser = DomainError{
		Code:    "USR-400101",
		Message: "Failed to create user",
	}
)

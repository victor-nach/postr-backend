package handlers


import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Users
type createUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zipcode   string `json:"zipcode"`
}

func (r createUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Firstname, validation.Required, validation.Length(2, 0)),
		validation.Field(&r.Lastname, validation.Required, validation.Length(2, 0)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Street, validation.Required),
		validation.Field(&r.City, validation.Required),
		validation.Field(&r.State, validation.Required),
		validation.Field(&r.Zipcode, validation.Required),
	)
}

// Posts
type createPostRequest struct {
	UserID string `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (r createPostRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required),
		validation.Field(&r.Title, validation.Required),
		validation.Field(&r.Body, validation.Required),
	)
}


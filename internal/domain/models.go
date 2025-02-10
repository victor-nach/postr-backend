package domain

import (
	"time"
)

type (
	User struct {
		ID        string    `json:"id"`
		Firstname string    `json:"firstname"`
		Lastname  string    `json:"lastname"`
		Email     string    `json:"email"`
		Street    string    `json:"street"`
		City      string    `json:"city"`
		State     string    `json:"state"`
		Zipcode   string    `json:"zipcode"`
		CreatedAt time.Time `json:"createdAt"`
	}

	Post struct {
		ID        string    `json:"id"`
		UserID    string    `json:"userId"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"createdAt"`
	}

	PaginatedUsers struct {
		Pagination Pagination `json:"pagination"`
		Users      []User     `json:"users"`
	}

	Pagination struct {
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
		TotalSize   int `json:"total_size"`
	}
)

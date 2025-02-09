package domain

import "time"

type User struct {
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

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

type PaginatedUsers struct {
	CurrentPage int    `json:"current_page"`
	TotalPages  int    `json:"total_pages"`
	TotalSize   int    `json:"total_size"`
	Users       []User `json:"users"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	List(pageNumber int, pageSize int) (PaginatedUsers, error)
	Count() (int, error)
	Validate(userID string) error
}

type PostRepository interface {
	Create(post *Post) error
	Delete(id string) error
	ListByUserID(userId string) ([]Post, error)
}

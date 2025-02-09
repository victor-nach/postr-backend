package handlers

import "github.com/victor-nach/postr-backend/internal/domain"

const (
	successStatus = "success"
	errorStatus   = "error"
)

type APIResponse struct {
	Status     string             `json:"status"`
	Message    string             `json:"message"`
	Pagination *domain.Pagination `json:"pagination,omitempty"`
	Data       any                `json:"data"`
}

type Count struct {
	Count int `json:"count"`
}

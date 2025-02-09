package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type UserHandler struct {
	repo   domain.UserRepository
	logger *zap.Logger
}

func NewUserHandler(repo domain.UserRepository, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		repo:   repo,
		logger: logger,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := req.Validate(); err != nil {
		if verrs, ok := err.(validation.Errors); ok {
			h.logger.Error("Validation errors", zap.Any("errors", verrs))
			c.JSON(http.StatusBadRequest, domain.ErrInvalidInput.WithFieldErrors(verrs))
			return
		}
		
		h.logger.Error("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrInvalidInput)
		return
	}

	user := &domain.User{
		ID:        uuid.New().String(),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Street:    req.Street,
		City:      req.City,
		State:     req.State,
		Zipcode:   req.Zipcode,
		CreatedAt: time.Now(),
	}

	if err := h.repo.Create(user); err != nil {
		h.logger.Error("Error creating user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.ErrCreateUser)
		return
	}

	h.logger.Info("User created successfully", zap.Any("user", user))
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	pageNumber, err := strconv.Atoi(c.Query("pageNumber"))
	if err != nil {
		pageNumber = 0 // default
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 10 // default
	}

	paginated, err := h.repo.List(pageNumber, pageSize)
	if err != nil {
		h.logger.Error("Error listing users", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServer)
		return
	}

	h.logger.Info("Users listed successfully", zap.Any("paginated", paginated))
	c.JSON(http.StatusOK, paginated)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Info("User not found", zap.String("id", id))
			c.JSON(http.StatusNotFound, domain.ErrUserNotFound)
			return
		}
		h.logger.Error("Error retrieving user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServer)
		return
	}

	h.logger.Info("User retrieved successfully", zap.Any("user", user))
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CountUsers(c *gin.Context) {
	count, err := h.repo.Count()
	if err != nil {
		h.logger.Error("Error counting users", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServer)
		return
	}
	h.logger.Info("Users count retrieved successfully", zap.Int("count", count))
	c.JSON(http.StatusOK, gin.H{"count": count})
}

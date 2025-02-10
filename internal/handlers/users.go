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

	"github.com/victor-nach/postr-backend/internal/domain"
)

type UserHandler struct {
	service domain.UserService
	logger  *zap.Logger
}

func NewUserHandler(service domain.UserService, logger *zap.Logger) *UserHandler {
	logger = logger.With(zap.String("package", "handlers"))

	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "CreateUser"))

	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logr.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate request body
	if err := req.Validate(); err != nil {
		if verrs, ok := err.(validation.Errors); ok {
			logr.Error("Validation errors", zap.Any("errors", verrs))
			c.JSON(http.StatusBadRequest, domain.ErrInvalidInput.WithFieldErrors(verrs))
			return
		}

		logr.Error("Validation error", zap.Error(err))
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

	if err := h.service.Create(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("User created successfully", zap.Any("user", user))

	resp := APIResponse{
		Status:  successStatus,
		Message: "Users created successfully",
		Data:    user,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "ListUsers"))

	pageNumber, err := strconv.Atoi(c.Query("pageNumber"))
	if err != nil {
		pageNumber = 1 // default
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 10 // default
	}

	paginatedUsers, err := h.service.List(c.Request.Context(), pageNumber, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("Users listed successfully", zap.Any("paginated", paginatedUsers))

	resp := APIResponse{
		Status:     successStatus,
		Message:    "Users listed successfully",
		Pagination: &paginatedUsers.Pagination,
		Data:       paginatedUsers.Users,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "GetUserByID"))

	id := c.Param("id")
	user, err := h.service.Get(c.Request.Context(), id)

	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("User retrieved successfully", zap.Any("user", user))

	resp := APIResponse{
		Status:  successStatus,
		Message: "User retrieved successfully",
		Data:    user,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) CountUsers(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "CountUsers"))

	count, err := h.service.Count(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("Users count retrieved successfully", zap.Int("count", count))

	resp := APIResponse{
		Status:  successStatus,
		Message: "Users count retrieved successfully",
		Data: Count{
			Count: count,
		},
	}
	c.JSON(http.StatusOK, resp)
}

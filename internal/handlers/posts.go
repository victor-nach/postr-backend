package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type PostHandler struct {
	service domain.PostService
	logger  *zap.Logger
}

func NewPostHandler(service domain.PostService, logger *zap.Logger) *PostHandler {
	logger = logger.With(zap.String("package", "handlers"))

	return &PostHandler{
		service: service,
		logger:  logger,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "CreatePost"))

	var req createPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logr.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrInvalidInput)
		return
	}

	// Trim whitespace from the request fields
	req.UserID = strings.TrimSpace(req.UserID)
	req.Title = strings.TrimSpace(req.Title)
	req.Body = strings.TrimSpace(req.Body)

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

	post := &domain.Post{
		ID:        uuid.NewString(),
		UserID:    req.UserID,
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
	}

	if err := h.service.Create(c.Request.Context(), post); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("Post created successfully", zap.Any("post", post))

	resp := APIResponse{
		Status:  successStatus,
		Message: "Posts listed successfully",
		Data:    post,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *PostHandler) ListPostsByUserID(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "ListPostsByUserID"))

	userId := c.Param("userId")
	if userId == "" {
		h.logger.Error("Missing userId path parameter")
		c.JSON(http.StatusBadRequest, domain.ErrInvalidInput)
		return
	}

	posts, err := h.service.List(c.Request.Context(), userId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusInternalServerError, domain.ErrInternalServer)
		return
	}

	logr.Info("Posts listed successfully", zap.String("userId", userId), zap.Int("count", len(posts)))

	resp := APIResponse{
		Status:  successStatus,
		Message: "Posts listed successfully",
		Data:    posts,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "ListPostsByUserID"))

	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, domain.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return

	}

	logr.Info("Post deleted successfully", zap.String("id", id))
	c.Status(http.StatusNoContent)
}

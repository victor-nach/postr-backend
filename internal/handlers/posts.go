package handlers

import (
	"errors"
	"net/http"
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
	return &PostHandler{
		service: service,
		logger:  logger,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req createPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrInvalidInput)
		return
	}

	// Validate request body
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

	h.logger.Info("Post created successfully", zap.Any("post", post))

	resp := APIResponse{
		Status:  successStatus,
		Message: "Posts listed successfully",
		Data:    post,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *PostHandler) ListPostsByUserID(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		h.logger.Error("Missing userId query parameter")
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

	h.logger.Info("Posts listed successfully", zap.String("userId", userId), zap.Int("count", len(posts)))

	resp := APIResponse{
		Status:  successStatus,
		Message: "Posts listed successfully",
		Data:    posts,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, domain.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return

	}

	h.logger.Info("Post deleted successfully", zap.String("id", id))
	c.Status(http.StatusNoContent)
}

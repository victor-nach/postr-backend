package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type PostHandler struct {
	repo     domain.PostRepository
	userRepo domain.UserRepository
	logger   *zap.Logger
}

func NewPostHandler(repo domain.PostRepository, userRepo domain.UserRepository, logger *zap.Logger) *PostHandler {
	return &PostHandler{
		repo:   repo,
		userRepo: userRepo,
		logger: logger,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req createPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrInvalidInput)
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

	if err := h.userRepo.Validate(req.UserID); err != nil {
		h.logger.Error("Invalid userID", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrUserNotFound)
		return
	}

	post := &domain.Post{
		ID:        uuid.NewString(),
		UserID:    req.UserID,
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
	}

	if err := h.repo.Create(post); err != nil {
		h.logger.Error("Error creating post", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServer)
		return
	}

	h.logger.Info("Post created successfully", zap.Any("post", post))
	c.JSON(http.StatusCreated, post)
	h.logger.Info("CreatePost handler completed")
}

func (h *PostHandler) ListPostsByUserID(c *gin.Context) {
	userId := c.Query("userId")
	if userId == "" {
		h.logger.Error("Missing userId query parameter")
		c.JSON(http.StatusBadRequest, domain.ErrInvalidInput)
		return
	}

	if err := h.userRepo.Validate(userId); err != nil {
		h.logger.Error("Invalid userID", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrUserNotFound)
		return
	}

	posts, err := h.repo.ListByUserID(userId)
	if err != nil {
		h.logger.Error("Error listing posts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServer)
		return
	}

	h.logger.Info("Posts listed successfully", zap.String("userId", userId), zap.Int("count", len(posts)))
	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Info("Post not found", zap.String("id", id))
			c.JSON(http.StatusNotFound, domain.ErrUserNotFound)
			return
		}
		h.logger.Error("Error deleting post", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServer)
		return

	}

	h.logger.Info("Post deleted successfully", zap.String("id", id))
	c.Status(http.StatusNoContent)
}

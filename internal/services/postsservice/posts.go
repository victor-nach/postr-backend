package postsservice

import (
	"context"
	"errors"
	
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type service struct {
	postsRepo postsRepo
	usersRepo usersRepo
	logger    *zap.Logger
}

func New(postsRepo postsRepo, usersRepo usersRepo, logger *zap.Logger) domain.PostService {
	return &service{
		usersRepo: usersRepo,
		postsRepo: postsRepo,
		logger:    logger,
	}
}

type postsRepo interface {
	Create(ctx context.Context, post *domain.Post) error
	ListByUserID(ctx context.Context, userId string) ([]domain.Post, error)
	Delete(ctx context.Context, id string) error
}

type usersRepo interface {
	Validate(ctx context.Context, userID string) error
}

func (h *service) Create(ctx context.Context, post *domain.Post) error {
	// Validate userID
	if err := h.validateUserID(ctx, post.UserID); err != nil {
		h.logger.Error("Invalid userID", zap.Error(err))
		return domain.ErrUserNotFound
	}

	if err := h.postsRepo.Create(ctx, post); err != nil {
		h.logger.Error("Error creating post", zap.Error(err))
		return domain.ErrInternalServer
	}

	h.logger.Info("Post created successfully", zap.Any("post", post))

	return nil
}

func (h *service) List(ctx context.Context, userID string) ([]domain.Post, error) {
	// Validate userID
	if err := h.validateUserID(ctx, userID); err != nil {
		h.logger.Error("Invalid userID", zap.Error(err))
		return []domain.Post{}, err
	}

	posts, err := h.postsRepo.ListByUserID(ctx, userID)
	if err != nil {
		h.logger.Error("Error listing posts", zap.Error(err))
		return []domain.Post{}, domain.ErrInternalServer
	}

	h.logger.Info("Posts listed successfully", zap.String("user_id", userID), zap.Int("count", len(posts)))
	return posts, nil
}

func (h *service) Delete(ctx context.Context, id string) error {
	if err := h.postsRepo.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Info("Post not found", zap.String("id", id))
			return domain.ErrUserNotFound
		}

		h.logger.Error("Error deleting post", zap.Error(err))
		return domain.ErrInternalServer

	}

	h.logger.Info("Post deleted successfully", zap.String("id", id))
	return nil
}

func (h *service) validateUserID(ctx context.Context, userID string) error{
	if err := h.usersRepo.Validate(ctx, userID); err != nil {
		return domain.ErrUserNotFound
	}

	return nil
}
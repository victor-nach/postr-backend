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
	logger = logger.With(zap.String("package", "postsservice"))

	return &service{
		usersRepo: usersRepo,
		postsRepo: postsRepo,
		logger:    logger,
	}
}


//go:generate mockgen -destination=./mocks/mock_postsrepo.go -package=mocks github.com/victor-nach/postr-backend/internal/services/postsservice postsRepo

type postsRepo interface {
	Create(ctx context.Context, post *domain.Post) error
	ListByUserID(ctx context.Context, userId string) ([]domain.Post, error)
	Delete(ctx context.Context, id string) error
}


//go:generate mockgen -destination=./mocks/mock_usersrepo.go -package=mocks github.com/victor-nach/postr-backend/internal/services/postsservice usersRepo
type usersRepo interface {
	Validate(ctx context.Context, userID string) error
}

func (h *service) Create(ctx context.Context, post *domain.Post) error {
	logr := h.logger.With(zap.String("method", "Create"))

	// Validate userID
	if err := h.validateUserID(ctx, post.UserID); err != nil {
		logr.Error("Invalid userID", zap.Error(err))
		return domain.ErrUserNotFound
	}

	if err := h.postsRepo.Create(ctx, post); err != nil {
		logr.Error("Error creating post", zap.Error(err))
		return domain.ErrInternalServer
	}

	logr.Info("Post created successfully", zap.Any("post", post))

	return nil
}

func (h *service) List(ctx context.Context, userID string) ([]domain.Post, error) {
	logr := h.logger.With(zap.String("method", "List"))
	
	// Validate userID
	if err := h.validateUserID(ctx, userID); err != nil {
		logr.Error("Invalid userID", zap.Error(err))
		return []domain.Post{}, err
	}

	posts, err := h.postsRepo.ListByUserID(ctx, userID)
	if err != nil {
		logr.Error("Error listing posts", zap.Error(err))
		return []domain.Post{}, domain.ErrInternalServer
	}

	logr.Info("Posts listed successfully", zap.String("user_id", userID), zap.Int("count", len(posts)))
	return posts, nil
}

func (h *service) Delete(ctx context.Context, id string) error {
	logr := h.logger.With(zap.String("method", "Delete"))

	if err := h.postsRepo.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logr.Info("Post not found", zap.String("id", id))
			return domain.ErrUserNotFound
		}

		logr.Error("Error deleting post", zap.Error(err))
		return domain.ErrInternalServer

	}

	logr.Info("Post deleted successfully", zap.String("id", id))
	return nil
}

func (h *service) validateUserID(ctx context.Context, userID string) error{
	if err := h.usersRepo.Validate(ctx, userID); err != nil {
		return domain.ErrUserNotFound
	}

	return nil
}
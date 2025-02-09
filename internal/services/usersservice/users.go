package usersservice

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type service struct {
	repo   usersRepo
	logger *zap.Logger
}

func New(repo usersRepo, logger *zap.Logger) domain.UserService {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

type usersRepo interface {
	Create(ctx context.Context, user *domain.User) error
	Get(ctx context.Context, id string) (*domain.User, error)
	List(ctx context.Context, pageNumber int, pageSize int) (domain.PaginatedUsers, error)
	Count(ctx context.Context, ) (int, error)
	Validate(ctx context.Context, userID string) error
}

func (h *service) Create(ctx context.Context, user *domain.User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()

	if err := h.repo.Create(ctx, user); err != nil {
		h.logger.Error("Error creating user", zap.Error(err))
		return err
	}

	h.logger.Info("User created successfully", zap.Any("user", user))

	return nil
}

func (h *service) Get(ctx context.Context, id string) (*domain.User, error) {
	user, err := h.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Info("User not found", zap.String("id", id))
			return nil, domain.ErrUserNotFound
		}
		
		h.logger.Error("Error retrieving user", zap.Error(err))
		return nil, domain.ErrInternalServer
	}

	h.logger.Info("User retrieved successfully", zap.Any("user", user))

	return user, nil
}

func (h *service) List(ctx context.Context, pageNumber int, pageSize int) (domain.PaginatedUsers, error) {
	paginatedUsers, err := h.repo.List(ctx, pageNumber, pageSize)
	if err != nil {
		h.logger.Error("Error listing users", zap.Error(err))
		return domain.PaginatedUsers{}, domain.ErrInternalServer
	}

	h.logger.Info("Users listed successfully", zap.Any("paginated", paginatedUsers))
	return paginatedUsers, nil
}

func (h *service) Count(ctx context.Context) (int, error) {
	count, err := h.repo.Count(ctx)
	if err != nil {
		h.logger.Error("Error counting users", zap.Error(err))
		return 0, domain.ErrInternalServer
	}

	h.logger.Info("Users count retrieved successfully", zap.Int("count", count))
	return count, nil
}
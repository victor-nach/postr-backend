package usersservice

import (
	"context"
	"errors"
	
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

//go:generate mockgen -destination=./mocks/mock_repo.go -package=mocks github.com/victor-nach/postr-backend/internal/services/usersservice usersRepo
type usersRepo interface {
	Create(ctx context.Context, user *domain.User) error
	Get(ctx context.Context, id string) (*domain.User, error)
	List(ctx context.Context, pageNumber int, pageSize int) (domain.PaginatedUsers, error)
	Count(ctx context.Context, ) (int, error)
	Validate(ctx context.Context, userID string) error
}

func (h *service) Create(ctx context.Context, user *domain.User) error {
	logr := h.logger.With(zap.String("method", "Create"))

	if err := h.repo.Create(ctx, user); err != nil {
		logr.Error("Error creating user", zap.Error(err))
		return err
	}

	logr.Info("User created successfully", zap.Any("user", user))

	return nil
}

func (h *service) Get(ctx context.Context, id string) (*domain.User, error) {
	logr := h.logger.With(zap.String("method", "Get"))

	user, err := h.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logr.Info("User not found", zap.String("id", id))
			return nil, domain.ErrUserNotFound
		}
		
		logr.Error("Error retrieving user", zap.Error(err))
		return nil, domain.ErrInternalServer
	}

	logr.Info("User retrieved successfully", zap.Any("user", user))

	return user, nil
}

func (h *service) List(ctx context.Context, pageNumber int, pageSize int) (domain.PaginatedUsers, error) {
	logr := h.logger.With(zap.String("method", "List"))

	paginatedUsers, err := h.repo.List(ctx, pageNumber, pageSize)
	if err != nil {
		logr.Error("Error listing users", zap.Error(err))
		return domain.PaginatedUsers{}, domain.ErrInternalServer
	}

	logr.Info("Users listed successfully", zap.Any("paginated", paginatedUsers))
	return paginatedUsers, nil
}

func (h *service) Count(ctx context.Context) (int, error) {
	logr := h.logger.With(zap.String("method", "Count"))

	count, err := h.repo.Count(ctx)
	if err != nil {
		logr.Error("Error counting users", zap.Error(err))
		return 0, domain.ErrInternalServer
	}

	logr.Info("Users count retrieved successfully", zap.Int("count", count))
	return count, nil
}
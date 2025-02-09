package repositories

import (
	"context"
	"math"

	"gorm.io/gorm"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Count(ctx context.Context) (int, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *userRepository) List(ctx context.Context, pageNumber int, pageSize int) (domain.PaginatedUsers, error) {
	var users []domain.User
	var total int64

	// Get total count of users
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Count(&total).Error; err != nil {
		return domain.PaginatedUsers{}, err
	}

	// Get paginated records
	offset := (pageNumber - 1) * pageSize
	if err := r.db.WithContext(ctx).Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return domain.PaginatedUsers{}, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	paginated := domain.PaginatedUsers{
		Pagination: domain.Pagination{
			CurrentPage: pageNumber,
			TotalPages:  totalPages,
			TotalSize:   int(total),
		},
		Users: users,
	}

	return paginated, nil
}

func (r *userRepository) Validate(ctx context.Context, userID string) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", userID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

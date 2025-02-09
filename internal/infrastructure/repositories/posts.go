package repositories

import (
	"context"

	"gorm.io/gorm"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(ctx context.Context, post *domain.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *postRepository) ListByUserID(ctx context.Context, userId string) ([]domain.Post, error) {
	var posts []domain.Post
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&domain.Post{}, "id = ?", id).Error
}
package repositories

import (
	"gorm.io/gorm"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) domain.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *domain.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) ListByUserID(userId string) ([]domain.Post, error) {
	var posts []domain.Post
	if err := r.db.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) Delete(id string) error {
	return r.db.Delete(&domain.Post{}, "id = ?", id).Error
}
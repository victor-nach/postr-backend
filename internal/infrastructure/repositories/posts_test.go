package repositories

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/victor-nach/postr-backend/internal/domain"

	_ "modernc.org/sqlite"
)

var (
	db      *gorm.DB
	sqlDB   *sql.DB
	postsrepo    *postRepository
	usersrepo    *userRepository
	testCtx = context.Background()
)

func TestMain(m *testing.M) {
	// Open an in-memory SQLite database
	var err error
	sqlDB, err = sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}

	db, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize GORM: %v", err)
	}

	// Apply migrations using gorm automigrate
	if err := db.AutoMigrate(&domain.User{}, &domain.Post{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	postsrepo = NewPostRepository(db)
	usersrepo = NewUserRepository(db)

	// Run the tests
	code := m.Run()

	sqlDB.Close()
	os.Exit(code)
}

func TestPostRepository_Create(t *testing.T) {
	post := domain.Post{
		ID:        uuid.NewString(),
		UserID:    uuid.NewString(),
		Title:     "Test Title",
		Body:      "Test Content body",
		CreatedAt: time.Now(),
	}

	err := postsrepo.Create(testCtx, &post)
	require.NoError(t, err)

	// Verify post creation
	var found domain.Post
	err = db.WithContext(testCtx).First(&found, "id = ?", post.ID).Error
	require.NoError(t, err)
	assert.Equal(t, post.Title, found.Title)
	assert.Equal(t, post.Body, found.Body)
}

func TestPostRepository_ListByUserID(t *testing.T) {
	posts := []domain.Post{
		{ID: uuid.NewString(), UserID: uuid.NewString(), Title: "Post 1", Body: "Body 1", CreatedAt: time.Now()},
		{ID: uuid.NewString(), UserID: uuid.NewString(), Title: "Post 2", Body: "Body 2", CreatedAt: time.Now()},
	}
	require.NoError(t, db.WithContext(testCtx).Create(&posts).Error)

	// List posts for the user
	result, err := postsrepo.ListByUserID(testCtx, posts[0].UserID)
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, "Post 1", result[0].Title)
}

func TestPostRepository_Delete(t *testing.T) {
	post := domain.Post{
		ID:        uuid.NewString(),
		UserID:    uuid.NewString(),
		Title:     "To Delete",
		Body:      "Body",
		CreatedAt: time.Now(),
	}
	require.NoError(t, db.WithContext(testCtx).Create(&post).Error)

	// Delete the post
	err := postsrepo.Delete(testCtx, post.ID)
	require.NoError(t, err)

	// Verify the post no longer exists
	var found domain.Post
	err = db.WithContext(testCtx).First(&found, "id = ?", post.ID).Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

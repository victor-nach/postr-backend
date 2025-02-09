package repositories

// import (
// 	// "context"
// 	// "database/sql"
// 	"fmt"
// 	// "log"
// 	// "os"
// 	"testing"

// 	// "gorm.io/driver/sqlite"
// 	"gorm.io/gorm"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"

// 	"github.com/victor-nach/postr-backend/internal/domain"
// )

// // cleanUsers deletes all user records to isolate tests.
// func cleanUsers(t *testing.T) {
// 	err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.User{}).Error
// 	require.NoError(t, err)
// }

// func TestUserRepository_Create(t *testing.T) {
// 	cleanUsers(t)

// 	user := domain.User{
// 		ID:    uuid.NewString(),
// 		Name:  "Test User",
// 		Email: "test@example.com",
// 	}

// 	err := repo.Create(testCtx, &user)
// 	require.NoError(t, err)

// 	var found domain.User
// 	err = db.WithContext(testCtx).First(&found, "id = ?", user.ID).Error
// 	require.NoError(t, err)
// 	assert.Equal(t, user.Name, found.Name)
// 	assert.Equal(t, user.Email, found.Email)
// }

// func TestUserRepository_Get(t *testing.T) {
// 	cleanUsers(t)

// 	user := domain.User{
// 		ID:    uuid.NewString(),
// 		Name:  "Get Test",
// 		Email: "get@example.com",
// 	}
// 	err := repo.Create(testCtx, &user)
// 	require.NoError(t, err)

// 	// Retrieve the user by ID.
// 	retrieved, err := repo.Get(testCtx, user.ID)
// 	require.NoError(t, err)
// 	assert.Equal(t, user.Name, retrieved.Name)
// 	assert.Equal(t, user.Email, retrieved.Email)

// 	// Attempt to get a non-existent user.
// 	_, err = repo.Get(testCtx, "non-existent-id")
// 	assert.Error(t, err)
// 	assert.Equal(t, gorm.ErrRecordNotFound, err)
// }

// func TestUserRepository_Count(t *testing.T) {
// 	cleanUsers(t)

// 	// Count should be zero initially.
// 	count, err := repo.Count(testCtx)
// 	require.NoError(t, err)
// 	assert.Equal(t, 0, count)

// 	// Create two users.
// 	users := []domain.User{
// 		{ID: uuid.NewString(), Name: "User 1", Email: "user1@example.com"},
// 		{ID: uuid.NewString(), Name: "User 2", Email: "user2@example.com"},
// 	}
// 	err = db.WithContext(testCtx).Create(&users).Error
// 	require.NoError(t, err)

// 	count, err = repo.Count(testCtx)
// 	require.NoError(t, err)
// 	assert.Equal(t, len(users), count)
// }

// func TestUserRepository_List(t *testing.T) {
// 	cleanUsers(t)

// 	// Create 5 users.
// 	var users []domain.User
// 	for i := 1; i <= 5; i++ {
// 		users = append(users, domain.User{
// 			ID:    uuid.NewString(),
// 			Name:  fmt.Sprintf("User %d", i),
// 			Email: fmt.Sprintf("user%d@example.com", i),
// 		})
// 	}
// 	err := db.WithContext(testCtx).Create(&users).Error
// 	require.NoError(t, err)

// 	// List page 1 with page size 2.
// 	paginated, err := repo.List(testCtx, 1, 2)
// 	require.NoError(t, err)
// 	assert.Equal(t, 1, paginated.Pagination.CurrentPage)
// 	// Total pages = ceil(5/2) = 3.
// 	assert.Equal(t, 3, paginated.Pagination.TotalPages)
// 	assert.Equal(t, 5, paginated.Pagination.TotalSize)
// 	assert.Len(t, paginated.Users, 2)

// 	// List page 2.
// 	paginated, err = repo.List(testCtx, 2, 2)
// 	require.NoError(t, err)
// 	assert.Equal(t, 2, paginated.Pagination.CurrentPage)
// 	assert.Len(t, paginated.Users, 2)

// 	// List page 3.
// 	paginated, err = repo.List(testCtx, 3, 2)
// 	require.NoError(t, err)
// 	assert.Equal(t, 3, paginated.Pagination.CurrentPage)
// 	assert.Len(t, paginated.Users, 1)

// 	// List page 4 should return 0 users.
// 	paginated, err = repo.List(testCtx, 4, 2)
// 	require.NoError(t, err)
// 	assert.Equal(t, 4, paginated.Pagination.CurrentPage)
// 	assert.Len(t, paginated.Users, 0)
// }

// func TestUserRepository_Validate(t *testing.T) {
// 	cleanUsers(t)

// 	user := domain.User{
// 		ID:    uuid.NewString(),
// 		Name:  "Validate User",
// 		Email: "validate@example.com",
// 	}
// 	err := repo.Create(testCtx, &user)
// 	require.NoError(t, err)

// 	// Validate an existing user.
// 	err = repo.Validate(testCtx, user.ID)
// 	require.NoError(t, err)

// 	// Validate a non-existent user should return the defined error.
// 	err = repo.Validate(testCtx, "non-existent-id")
// 	assert.Error(t, err)
// 	assert.Equal(t, domain.ErrUserNotFound, err)
// }

package repositories

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/victor-nach/postr-backend/internal/domain"
)


func TestUserRepository_Create(t *testing.T) {
	cleanUsers(t)

	user := domain.User{
		ID:        uuid.NewString(),
		Firstname: "Test",
		Lastname:  "User",
		Email:     "test@example.com",
		Street:    "123 Test St",
		City:      "Testville",
		State:     "TS",
		Zipcode:   "12345",
		CreatedAt: time.Now(),
	}

	err := usersrepo.Create(testCtx, &user)
	require.NoError(t, err)

	var found domain.User
	err = db.WithContext(testCtx).First(&found, "id = ?", user.ID).Error
	require.NoError(t, err)
	assert.Equal(t, user.Firstname, found.Firstname)
	assert.Equal(t, user.Lastname, found.Lastname)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Street, found.Street)
	assert.Equal(t, user.City, found.City)
	assert.Equal(t, user.State, found.State)
	assert.Equal(t, user.Zipcode, found.Zipcode)
}

func TestUserRepository_Get(t *testing.T) {
	cleanUsers(t)

	user := domain.User{
		ID:        uuid.NewString(),
		Firstname: "Get",
		Lastname:  "Test",
		Email:     "get@example.com",
		Street:    "456 Get St",
		City:      "Getville",
		State:     "GT",
		Zipcode:   "67890",
		CreatedAt: time.Now(),
	}
	err := usersrepo.Create(testCtx, &user)
	require.NoError(t, err)

	// Retrieve the user by ID
	retrieved, err := usersrepo.Get(testCtx, user.ID)
	require.NoError(t, err)
	assert.Equal(t, user.Firstname, retrieved.Firstname)
	assert.Equal(t, user.Lastname, retrieved.Lastname)
	assert.Equal(t, user.Email, retrieved.Email)
	assert.Equal(t, user.Street, retrieved.Street)
	assert.Equal(t, user.City, retrieved.City)
	assert.Equal(t, user.State, retrieved.State)
	assert.Equal(t, user.Zipcode, retrieved.Zipcode)

	// Non-existent user
	_, err = usersrepo.Get(testCtx, "non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUserRepository_Count(t *testing.T) {
	cleanUsers(t)

	count, err := usersrepo.Count(testCtx)
	require.NoError(t, err)
	assert.Equal(t, 0, count)

	users := []domain.User{
		{
			ID:        uuid.NewString(),
			Firstname: "User",
			Lastname:  "One",
			Email:     "user1@example.com",
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.NewString(),
			Firstname: "User",
			Lastname:  "Two",
			Email:     "user2@example.com",
			CreatedAt: time.Now(),
		},
	}
	err = db.WithContext(testCtx).Create(&users).Error
	require.NoError(t, err)

	count, err = usersrepo.Count(testCtx)
	require.NoError(t, err)
	assert.Equal(t, len(users), count)
}

func TestUserRepository_List(t *testing.T) {
	cleanUsers(t)

	var users []domain.User
	for i := 1; i <= 5; i++ {
		users = append(users, domain.User{
			ID:        uuid.NewString(),
			Firstname: fmt.Sprintf("User%d", i),
			Lastname:  "Test",
			Email:     fmt.Sprintf("user%d@example.com", i),
			CreatedAt: time.Now(),
		})
	}
	err := db.WithContext(testCtx).Create(&users).Error
	require.NoError(t, err)

	// List page 1 with page size 2
	paginated, err := usersrepo.List(testCtx, 1, 2)
	require.NoError(t, err)
	assert.Equal(t, 1, paginated.Pagination.CurrentPage)
	assert.Equal(t, 3, paginated.Pagination.TotalPages)
	assert.Equal(t, 5, paginated.Pagination.TotalSize)
	assert.Len(t, paginated.Users, 2)

	// List page 2
	paginated, err = usersrepo.List(testCtx, 2, 2)
	require.NoError(t, err)
	assert.Equal(t, 2, paginated.Pagination.CurrentPage)
	assert.Len(t, paginated.Users, 2)
}

func TestUserRepository_Validate(t *testing.T) {
	cleanUsers(t)

	user := domain.User{
		ID:        uuid.NewString(),
		Firstname: "Validate",
		Lastname:  "User",
		Email:     "validate@example.com",
		CreatedAt: time.Now(),
	}
	err := usersrepo.Create(testCtx, &user)
	require.NoError(t, err)

	// Validate an existing user
	err = usersrepo.Validate(testCtx, user.ID)
	require.NoError(t, err)

	// Validate a non-existent user
	err = usersrepo.Validate(testCtx, "non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrUserNotFound, err)
}

func cleanUsers(t *testing.T) {
	err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.User{}).Error
	require.NoError(t, err)
}
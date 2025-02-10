package postsservice_test

import (
	"context"
	"testing"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/victor-nach/postr-backend/internal/domain"
	"github.com/victor-nach/postr-backend/internal/services/postsservice"
	"github.com/victor-nach/postr-backend/internal/services/postsservice/mocks"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostsRepo := mocks.NewMockpostsRepo(ctrl)
	mockUsersRepo := mocks.NewMockusersRepo(ctrl)

	logger := zap.NewNop()
	svc := postsservice.New(mockPostsRepo, mockUsersRepo, logger)

	ctx := context.Background()
	post := &domain.Post{
		ID:     uuid.NewString(),
		UserID: uuid.NewString(),
		Title: "Title 1",
		CreatedAt: time.Now(),
	}

	mockUsersRepo.EXPECT().Validate(ctx, post.UserID).Return(nil)
	mockPostsRepo.EXPECT().Create(ctx, post).Return(nil)

	err := svc.Create(ctx, post)
	require.NoError(t, err)
}

func TestService_List_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostsRepo := mocks.NewMockpostsRepo(ctrl)
	mockUsersRepo := mocks.NewMockusersRepo(ctrl)

	logger := zap.NewNop()
	svc := postsservice.New(mockPostsRepo, mockUsersRepo, logger)

	ctx := context.Background()
	userID := uuid.NewString()

	expectedPosts := []domain.Post{
		{
			ID:        uuid.NewString(),
			UserID:    userID,
			Title:     "Post 1",
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.NewString(),
			UserID:    userID,
			Title:     "Post 2",
			CreatedAt: time.Now(),
		},
	}

	mockUsersRepo.EXPECT().Validate(ctx, userID).Return(nil)
	mockPostsRepo.EXPECT().ListByUserID(ctx, userID).Return(expectedPosts, nil)

	posts, err := svc.List(ctx, userID)
	require.NoError(t, err)
	require.Equal(t, expectedPosts, posts)
}

func TestService_List_InvalidUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostsRepo := mocks.NewMockpostsRepo(ctrl)
	mockUsersRepo := mocks.NewMockusersRepo(ctrl)

	logger := zap.NewNop()
	svc := postsservice.New(mockPostsRepo, mockUsersRepo, logger)

	ctx := context.Background()
	userID := uuid.NewString()

	mockUsersRepo.EXPECT().Validate(ctx, userID).Return(domain.ErrUserNotFound)

	posts, err := svc.List(ctx, userID)
	require.Error(t, err)
	require.Equal(t, domain.ErrUserNotFound, err)
	require.Empty(t, posts)
}

func TestService_Delete_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostsRepo := mocks.NewMockpostsRepo(ctrl)
	mockUsersRepo := mocks.NewMockusersRepo(ctrl)

	logger := zap.NewNop()
	svc := postsservice.New(mockPostsRepo, mockUsersRepo, logger)

	ctx := context.Background()
	postID := uuid.NewString()

	mockPostsRepo.EXPECT().Delete(ctx, postID).Return(nil)

	err := svc.Delete(ctx, postID)
	require.NoError(t, err)
}

func TestService_Delete_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostsRepo := mocks.NewMockpostsRepo(ctrl)
	mockUsersRepo := mocks.NewMockusersRepo(ctrl)

	logger := zap.NewNop()
	svc := postsservice.New(mockPostsRepo, mockUsersRepo, logger)

	ctx := context.Background()
	postID := uuid.NewString()

	mockPostsRepo.EXPECT().Delete(ctx, postID).Return(gorm.ErrRecordNotFound)

	err := svc.Delete(ctx, postID)
	require.Error(t, err)
	require.Equal(t, domain.ErrUserNotFound, err)
}
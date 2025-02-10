package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/victor-nach/postr-backend/internal/domain"
	"github.com/victor-nach/postr-backend/internal/domain/mocks"
)

func TestPostHandler_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostService := mocks.NewMockPostService(ctrl)
	logger := zap.NewNop()
	handler := NewPostHandler(mockPostService, logger)

	reqBody := `{"userId": "b63df572-9bd1-4a4f-9f0d-2a8155a81fde", "title": "Test Title", "body": "Test Body"}`
	req, err := http.NewRequest("POST", "/posts", strings.NewReader(reqBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	mockPostService.EXPECT().Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, post *domain.Post) error {
			require.Equal(t, "b63df572-9bd1-4a4f-9f0d-2a8155a81fde", post.UserID)
			require.Equal(t, "Test Title", post.Title)
			require.Equal(t, "Test Body", post.Body)
			require.NotEmpty(t, post.ID)
			require.False(t, post.CreatedAt.IsZero())
			return nil
		}).Times(1)

	handler.CreatePost(c)

	require.Equal(t, http.StatusOK, w.Code)

	var resp APIResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "success", resp.Status)
	require.Equal(t, "Posts listed successfully", resp.Message)

	data, ok := resp.Data.(map[string]interface{})
	require.True(t, ok, "expected Data to be a map")
	require.Equal(t, "b63df572-9bd1-4a4f-9f0d-2a8155a81fde", data["userId"])
	require.Equal(t, "Test Title", data["title"])
	require.Equal(t, "Test Body", data["body"])
}

func TestPostHandler_ListPostsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostService := mocks.NewMockPostService(ctrl)
	logger := zap.NewNop()
	handler := NewPostHandler(mockPostService, logger)

	req, err := http.NewRequest("GET", "/posts/b63df572-9bd1-4a4f-9f0d-2a8155a81fde", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "userId", Value: "b63df572-9bd1-4a4f-9f0d-2a8155a81fde"}}

	expectedPosts := []domain.Post{
		{
			ID:        "post1",
			UserID:    "12345",
			Title:     "Title 1",
			Body:      "Body 1",
			CreatedAt: time.Now(),
		},
		{
			ID:        "post2",
			UserID:    "12345",
			Title:     "Title 2",
			Body:      "Body 2",
			CreatedAt: time.Now(),
		},
	}

	mockPostService.EXPECT().List(gomock.Any(), "b63df572-9bd1-4a4f-9f0d-2a8155a81fde").Return(expectedPosts, nil).Times(1)

	handler.ListPostsByUserID(c)

	require.Equal(t, http.StatusOK, w.Code)

	var resp APIResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "success", resp.Status)
	require.Equal(t, "Posts listed successfully", resp.Message)

	dataSlice, ok := resp.Data.([]interface{})
	require.True(t, ok, "expected Data to be a slice")
	require.Len(t, dataSlice, len(expectedPosts))
}

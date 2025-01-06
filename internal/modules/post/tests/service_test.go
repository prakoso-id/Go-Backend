package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-backend/internal/modules/post/domain/entity"
	"go-backend/internal/modules/post/domain/service"
	"go-backend/internal/modules/post/dto"
	"go-backend/internal/modules/post/mocks"
)

func TestCreatePostService(t *testing.T) {
	mockRepo := new(mocks.MockPostRepository)
	svc := service.NewPostService(mockRepo)

	tests := []struct {
		name          string
		userID        uint
		input         *dto.CreatePostRequest
		setupMock     func()
		expectedResp  *dto.CreatePostResponse
		expectedError error
	}{
		{
			name:   "Success",
			userID: 1,
			input: &dto.CreatePostRequest{
				Title:   "Test Post",
				Content: "Test Content",
			},
			setupMock: func() {
				mockRepo.On("Create", mock.MatchedBy(func(post *entity.Post) bool {
					return post.Title == "Test Post" &&
						post.Content == "Test Content" &&
						post.UserID == uint(1)
				})).Run(func(args mock.Arguments) {
					post := args.Get(0).(*entity.Post)
					post.ID = 1
				}).Return(nil)
			},
			expectedResp: &dto.CreatePostResponse{
				ID:      1,
				Title:   "Test Post",
				Content: "Test Content",
				UserID:  1,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			resp, err := svc.Create(tt.userID, tt.input)
			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedError, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
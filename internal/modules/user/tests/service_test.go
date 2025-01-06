package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-backend/internal/modules/user/domain/service"
	"go-backend/internal/modules/user/dto"
	"go-backend/internal/modules/user/mocks"
)

func TestUserService_Create(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	svc := service.NewUserService(mockRepo)

	tests := []struct {
		name    string
		input   *dto.CreateUserRequest
		mockFn  func()
		wantErr bool
	}{
		{
			name: "success",
			input: &dto.CreateUserRequest{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password123",
			},
			mockFn: func() {
				mockRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			user, err := svc.Create(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, user)
			mockRepo.AssertExpectations(t)
		})
	}
}
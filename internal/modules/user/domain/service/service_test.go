package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"go-backend/internal/modules/user/domain/entity"
	"go-backend/internal/modules/user/dto"
	"go-backend/internal/modules/user/mocks"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

func TestUserService_Create(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	svc := NewUserService(mockRepo)

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

func TestUserService_GetByID(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	svc := NewUserService(mockRepo)

	tests := []struct {
		name    string
		id      uint
		mockFn  func()
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			mockFn: func() {
				mockRepo.On("GetByID", uint(1)).Return(&entity.User{
					ID:    1,
					Name:  "Test User",
					Email: "test@example.com",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "not found",
			id:   2,
			mockFn: func() {
				mockRepo.On("GetByID", uint(2)).Return(nil, ErrRecordNotFound)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			user, err := svc.GetByID(tt.id)
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

func TestUserService_GetByEmail(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	svc := NewUserService(mockRepo)

	tests := []struct {
		name    string
		email   string
		mockFn  func()
		wantErr bool
	}{
		{
			name:  "success",
			email: "test@example.com",
			mockFn: func() {
				mockRepo.On("GetByEmail", "test@example.com").Return(&entity.User{
					ID:    1,
					Name:  "Test User",
					Email: "test@example.com",
				}, nil)
			},
			wantErr: false,
		},
		{
			name:  "not found",
			email: "notfound@example.com",
			mockFn: func() {
				mockRepo.On("GetByEmail", "notfound@example.com").Return(nil, ErrRecordNotFound)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			user, err := svc.GetByEmail(tt.email)
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

func TestUserService_Update(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	svc := NewUserService(mockRepo)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("newpassword123"), bcrypt.DefaultCost)

	tests := []struct {
		name    string
		id      uint
		input   *dto.UpdateUserRequest
		mockFn  func()
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			input: &dto.UpdateUserRequest{
				Name:     "Updated User",
				Password: "newpassword123",
			},
			mockFn: func() {
				mockRepo.On("GetByID", uint(1)).Return(&entity.User{
					ID:       1,
					Name:     "Test User",
					Email:    "test@example.com",
					Password: string(hashedPassword),
				}, nil)
				mockRepo.On("Update", mock.AnythingOfType("*entity.User")).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			user, err := svc.Update(tt.id, tt.input)
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

func TestUserService_Delete(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	svc := NewUserService(mockRepo)

	tests := []struct {
		name    string
		id      uint
		mockFn  func()
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			mockFn: func() {
				mockRepo.On("Delete", uint(1)).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			err := svc.Delete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUserService_List(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	svc := NewUserService(mockRepo)

	tests := []struct {
		name    string
		mockFn  func()
		wantErr bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockRepo.On("List", 1, 10).Return([]*entity.User{
					{
						ID:    1,
						Name:  "User 1",
						Email: "user1@example.com",
					},
					{
						ID:    2,
						Name:  "User 2",
						Email: "user2@example.com",
					},
				}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			users, err := svc.List(1, 10)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, users)
			assert.Len(t, users, 2)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUserService_Login(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	svc := NewUserService(mockRepo)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	tests := []struct {
		name    string
		input   *dto.LoginRequest
		mockFn  func()
		wantErr bool
	}{
		{
			name: "success",
			input: &dto.LoginRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			mockFn: func() {
				mockRepo.On("GetByEmail", "test@example.com").Return(&entity.User{
					ID:       1,
					Email:    "test@example.com",
					Password: string(hashedPassword),
				}, nil)
				mockRepo.On("Update", mock.AnythingOfType("*entity.User")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "invalid credentials",
			input: &dto.LoginRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			mockFn: func() {
				mockRepo.On("GetByEmail", "test@example.com").Return(&entity.User{
					ID:       1,
					Email:    "test@example.com",
					Password: string(hashedPassword),
				}, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			resp, err := svc.Login(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Token)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUserService_Logout(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	svc := NewUserService(mockRepo)

	tests := []struct {
		name    string
		userID  uint
		mockFn  func()
		wantErr bool
	}{
		{
			name:   "success",
			userID: 1,
			mockFn: func() {
				mockRepo.On("GetByID", uint(1)).Return(&entity.User{
					ID:    1,
					Email: "test@example.com",
				}, nil)
				mockRepo.On("Update", mock.AnythingOfType("*entity.User")).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			err := svc.Logout(tt.userID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

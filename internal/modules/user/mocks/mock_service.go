package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-backend/internal/modules/user/dto"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserResponse), args.Error(1)
}

func (m *MockUserService) GetByID(id uint) (*dto.UserResponse, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserResponse), args.Error(1)
}

func (m *MockUserService) GetByEmail(email string) (*dto.UserResponse, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserResponse), args.Error(1)
}

func (m *MockUserService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.LoginResponse), args.Error(1)
}

func (m *MockUserService) Update(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	args := m.Called(id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserResponse), args.Error(1)
}

func (m *MockUserService) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserService) List(page, limit int) ([]*dto.UserResponse, error) {
	args := m.Called(page, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*dto.UserResponse), args.Error(1)
}

func (m *MockUserService) Logout(userID uint) error {
	args := m.Called(userID)
	return args.Error(0)
}

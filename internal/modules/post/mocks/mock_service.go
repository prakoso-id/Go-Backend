package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-backend/internal/modules/post/dto"
)

type MockPostService struct {
	mock.Mock
}

func (m *MockPostService) Create(userID uint, req *dto.CreatePostRequest) (*dto.CreatePostResponse, error) {
	args := m.Called(userID, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CreatePostResponse), args.Error(1)
}

func (m *MockPostService) GetByID(id uint) (*dto.GetPostResponse, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.GetPostResponse), args.Error(1)
}

func (m *MockPostService) Update(id, userID uint, req *dto.UpdatePostRequest) (*dto.UpdatePostResponse, error) {
	args := m.Called(id, userID, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UpdatePostResponse), args.Error(1)
}

func (m *MockPostService) Delete(id, userID uint) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

func (m *MockPostService) List(page, pageSize int) ([]dto.GetPostResponse, error) {
	args := m.Called(page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dto.GetPostResponse), args.Error(1)
}

func (m *MockPostService) ListByUserID(userID uint, page, pageSize int) ([]dto.GetPostResponse, error) {
	args := m.Called(userID, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dto.GetPostResponse), args.Error(1)
}

package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-backend/internal/modules/post/domain/entity"
)

type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) Create(post *entity.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func (m *MockPostRepository) GetByID(id uint) (*entity.Post, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Post), args.Error(1)
}

func (m *MockPostRepository) Update(post *entity.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func (m *MockPostRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPostRepository) List(offset, limit int) ([]entity.Post, error) {
	args := m.Called(offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Post), args.Error(1)
}

func (m *MockPostRepository) ListByUserID(userID uint, offset, limit int) ([]entity.Post, error) {
	args := m.Called(userID, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Post), args.Error(1)
}

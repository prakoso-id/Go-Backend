package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-backend/internal/modules/project/domain/entity"
)

type MockProjectRepository struct {
	mock.Mock
}

func (m *MockProjectRepository) Create(project *entity.Project) error {
	args := m.Called(project)
	return args.Error(0)
}

func (m *MockProjectRepository) GetByID(id uint) (*entity.Project, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Project), args.Error(1)
}

func (m *MockProjectRepository) GetAll() ([]entity.Project, error) {
	args := m.Called()
	return args.Get(0).([]entity.Project), args.Error(1)
}

func (m *MockProjectRepository) Update(project *entity.Project) error {
	args := m.Called(project)
	return args.Error(0)
}

func (m *MockProjectRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProjectRepository) GetByUserID(userID uint) ([]entity.Project, error) {
	args := m.Called(userID)
	return args.Get(0).([]entity.Project), args.Error(1)
}
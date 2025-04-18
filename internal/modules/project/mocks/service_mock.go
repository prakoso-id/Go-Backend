package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-backend/internal/modules/project/domain/entity"
	"go-backend/internal/modules/project/dto"
)

type MockProjectService struct {
	mock.Mock
}

func (m *MockProjectService) Create(project *entity.Project) (*dto.CreateProjectResponse, error) {
	args := m.Called(project)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CreateProjectResponse), args.Error(1)
}

func (m *MockProjectService) GetByID(id uint) (*dto.ProjectResponse, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.ProjectResponse), args.Error(1)
}

func (m *MockProjectService) GetAll() ([]dto.ProjectResponse, error) {
	args := m.Called()
	return args.Get(0).([]dto.ProjectResponse), args.Error(1)
}

func (m *MockProjectService) Update(id uint, userID uint, req *dto.UpdateProjectRequest) (*dto.UpdateProjectResponse, error) {
	args := m.Called(id, userID, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UpdateProjectResponse), args.Error(1)
}

func (m *MockProjectService) Delete(id, userID uint) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

func (m *MockProjectService) GetByUserID(userID uint) ([]dto.ProjectResponse, error) {
	args := m.Called(userID)
	return args.Get(0).([]dto.ProjectResponse), args.Error(1)
}
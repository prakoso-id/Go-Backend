package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-backend/internal/modules/project/domain/entity"
	"go-backend/internal/modules/project/domain/service"
	"go-backend/internal/modules/project/dto"
	"go-backend/internal/modules/project/mocks"
)

func TestCreateProjectService(t *testing.T) {
	mockRepo := new(mocks.MockProjectRepository)
	svc := service.NewProjectService(mockRepo)

	tests := []struct {
		name             string
		input            *entity.Project
		setupMock        func()
		expectedResponse *dto.CreateProjectResponse
		expectedError    error
	}{
		{
			name: "Success",
			input: &entity.Project{
				Name:        "Test Project",
				Description: "Test Description",
			},
			setupMock: func() {
				mockRepo.On("Create", mock.AnythingOfType("*entity.Project")).
					Return(nil)
			},
			expectedResponse: &dto.CreateProjectResponse{
				ID:          0,
				Name:        "Test Project",
				Description: "Test Description",
				UserID:      0,
			},
			expectedError: nil,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			resp, err := svc.Create(tt.input)
			assert.Equal(t, tt.expectedError, err)
			if tt.expectedError == nil {
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectedResponse.Name, resp.Name)
				assert.Equal(t, tt.expectedResponse.Description, resp.Description)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
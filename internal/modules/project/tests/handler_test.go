package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-backend/internal/modules/project/dto"
	"go-backend/internal/modules/project/handlers"
	"go-backend/internal/modules/project/mocks"
)

func TestCreateProject(t *testing.T) {
	mockService := new(mocks.MockProjectService)
	handler := handlers.NewProjectHandler(mockService, nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/projects", handler.Create)

	tests := []struct {
		name          string
		input         dto.CreateProjectRequest
		setupMock     func()
		expectedCode  int
		expectedError string
	}{
		{
			name: "Success",
			input: dto.CreateProjectRequest{
				Name:        "Test Project",
				Description: "Test Description",
			},
			setupMock: func() {
				mockService.On("Create", mock.AnythingOfType("*entity.Project")).
					Return(&dto.CreateProjectResponse{
						ID:          1,
						Name:        "Test Project",
						Description: "Test Description",
						UserID:      1,
					}, nil)
			},
			expectedCode: http.StatusCreated,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			body, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}

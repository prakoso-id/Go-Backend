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
	"go-backend/internal/modules/user/dto"
	"go-backend/internal/modules/user/handlers"
	"go-backend/internal/modules/user/mocks"
)

func TestUserHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockUserService)
	handler := handlers.NewUserHandler(mockService)

	tests := []struct {
		name           string
		input          *dto.CreateUserRequest
		mockFn         func()
		expectedStatus int
	}{
		{
			name: "success",
			input: &dto.CreateUserRequest{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "password123",
			},
			mockFn: func() {
				mockService.On("Create", mock.AnythingOfType("*dto.CreateUserRequest")).
					Return(&dto.UserResponse{
						ID:    1,
						Name:  "Test User",
						Email: "test@example.com",
					}, nil)
			},
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonData, _ := json.Marshal(tt.input)
			c.Request = httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
			c.Request.Header.Set("Content-Type", "application/json")

			tt.mockFn()
			handler.Create(c)

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}
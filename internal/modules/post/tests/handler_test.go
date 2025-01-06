package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"go-backend/internal/modules/post/dto"
	"go-backend/internal/modules/post/handlers"
	"go-backend/internal/modules/post/mocks"
)

func setupTest() (*gin.Engine, *mocks.MockPostService) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(mocks.MockPostService)
	handler := handlers.NewPostHandler(mockService)

	// Add auth middleware before routes
	r.Use(func(c *gin.Context) {
		c.Set("user_id", uint(1))
		c.Next()
	})

	group := r.Group("/posts")
	group.POST("", handler.Create)
	group.GET("", handler.List)
	group.GET("/:id", handler.GetByID)
	group.PUT("/:id", handler.Update)
	group.DELETE("/:id", handler.Delete)

	return r, mockService
}

func TestPostHandler_Create(t *testing.T) {
	r, mockService := setupTest()

	t.Run("Success", func(t *testing.T) {
		req := dto.CreatePostRequest{
			Title:   "Test Post",
			Content: "Test Content",
		}

		response := &dto.CreatePostResponse{
			ID:      1,
			Title:   req.Title,
			Content: req.Content,
			UserID:  1,
		}

		mockService.On("Create", uint(1), &req).Return(response, nil).Once()

		body, _ := json.Marshal(req)
		w := httptest.NewRecorder()
		req1, _ := http.NewRequest(http.MethodPost, "/posts", bytes.NewBuffer(body))

		r.ServeHTTP(w, req1)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Invalid Request", func(t *testing.T) {
		req := dto.CreatePostRequest{
			Content: "Test Content",
		}

		body, _ := json.Marshal(req)
		w := httptest.NewRecorder()
		req1, _ := http.NewRequest(http.MethodPost, "/posts", bytes.NewBuffer(body))

		r.ServeHTTP(w, req1)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Service Error", func(t *testing.T) {
		req := dto.CreatePostRequest{
			Title:   "Test Post",
			Content: "Test Content",
		}

		mockService.On("Create", uint(1), &req).Return(nil, errors.New("service error")).Once()

		body, _ := json.Marshal(req)
		w := httptest.NewRecorder()
		req1, _ := http.NewRequest(http.MethodPost, "/posts", bytes.NewBuffer(body))

		r.ServeHTTP(w, req1)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})
}
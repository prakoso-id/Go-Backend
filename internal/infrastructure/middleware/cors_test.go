package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSMiddleware(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CORSMiddleware())
	
	// Add a test endpoint
	router.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	
	// Test cases
	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "GET request",
			method:         "GET",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "OPTIONS request",
			method:         "OPTIONS",
			expectedStatus: http.StatusNoContent,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, "/test", nil)
			router.ServeHTTP(w, req)
			
			// Assert response status
			assert.Equal(t, tt.expectedStatus, w.Code)
			
			// Assert CORS headers
			assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
			assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
			assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Content-Type")
			assert.Contains(t, w.Header().Get("Access-Control-Allow-Methods"), "POST")
		})
	}
}

package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-backend/internal/infrastructure/database"
	"go-backend/internal/infrastructure/middleware"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	return router
}

func TestSetupRouter(t *testing.T) {
	db := database.SetupTestDB(t)
	defer database.CleanupTestDB(t, db)

	router := setupTestRouter()

	// Test CORS middleware
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	req.Header.Set("Access-Control-Request-Method", "POST")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "POST, OPTIONS, GET, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
}

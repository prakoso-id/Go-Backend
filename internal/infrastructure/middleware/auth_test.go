package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"go-backend/internal/infrastructure/database"
	"go-backend/internal/modules/user/domain/entity"
	"go-backend/internal/modules/user/domain/repository"
	"gorm.io/gorm"
)

func cleanupTestData(t *testing.T, db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error
	assert.NoError(t, err)
}

func setupTestRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	db := database.SetupTestDB(t)
	cleanupTestData(t, db)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	return router, db
}

func generateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func TestJWTAuth_ValidToken(t *testing.T) {
	router, db := setupTestRouter(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	// Create a test user with a valid token
	userRepo := repository.NewUserRepository(db)
	user := &entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	token, err := generateToken(user.ID)
	assert.NoError(t, err)

	expiresAt := time.Now().Add(24 * time.Hour)
	user.SetToken(token, expiresAt)
	err = userRepo.Update(user)
	assert.NoError(t, err)

	// Setup test endpoint
	router.GET("/test", JWTAuth(AccessToken), func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		assert.True(t, exists)
		assert.Equal(t, user.ID, userID.(uint))
		c.String(http.StatusOK, "success")
	})

	// Make request with valid token
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "success", w.Body.String())
}

func TestJWTAuth_InvalidToken(t *testing.T) {
	router, db := setupTestRouter(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	// Setup test endpoint
	router.GET("/test", JWTAuth(AccessToken), func(c *gin.Context) {
		t.Error("This handler should not be called")
	})

	// Make request with invalid token
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestJWTAuth_MissingToken(t *testing.T) {
	router, db := setupTestRouter(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	// Setup test endpoint
	router.GET("/test", JWTAuth(AccessToken), func(c *gin.Context) {
		t.Error("This handler should not be called")
	})

	// Make request without token
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestJWTAuth_ExpiredToken(t *testing.T) {
	router, db := setupTestRouter(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	// Create a test user with an expired token
	userRepo := repository.NewUserRepository(db)
	user := &entity.User{
		Name:     "Test User 2",
		Email:    "test2@example.com",
		Password: "password123",
	}
	err := userRepo.Create(user)
	assert.NoError(t, err)

	// Generate expired token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(-24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	assert.NoError(t, err)

	expiresAt := time.Now().Add(-24 * time.Hour)
	user.SetToken(tokenString, expiresAt)
	err = userRepo.Update(user)
	assert.NoError(t, err)

	// Setup test endpoint
	router.GET("/test", JWTAuth(AccessToken), func(c *gin.Context) {
		t.Error("This handler should not be called")
	})

	// Make request with expired token
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

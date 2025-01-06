package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go-backend/internal/infrastructure/database"
	"go-backend/internal/modules/user/domain/entity"
	"gorm.io/gorm"
)

func cleanupTestData(t *testing.T, db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error
	assert.NoError(t, err)
}

func TestUserRepository_Create(t *testing.T) {
	db := database.SetupTestDB(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	repo := NewUserRepository(db)

	user := &entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}

	err := repo.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
}

func TestUserRepository_GetByID(t *testing.T) {
	db := database.SetupTestDB(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	repo := NewUserRepository(db)

	// Create a test user
	user := &entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}
	err := repo.Create(user)
	assert.NoError(t, err)

	// Get the user by ID
	found, err := repo.GetByID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.ID, found.ID)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
}

func TestUserRepository_GetByID_NotFound(t *testing.T) {
	db := database.SetupTestDB(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	repo := NewUserRepository(db)

	found, err := repo.GetByID(999)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db := database.SetupTestDB(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	repo := NewUserRepository(db)

	// Create a test user
	user := &entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}
	err := repo.Create(user)
	assert.NoError(t, err)

	// Get the user by email
	found, err := repo.GetByEmail(user.Email)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.ID, found.ID)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
}

func TestUserRepository_GetByEmail_NotFound(t *testing.T) {
	db := database.SetupTestDB(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	repo := NewUserRepository(db)

	found, err := repo.GetByEmail("nonexistent@example.com")
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestUserRepository_Update(t *testing.T) {
	db := database.SetupTestDB(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	repo := NewUserRepository(db)

	// Create a test user
	user := &entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}
	err := repo.Create(user)
	assert.NoError(t, err)

	// Update the user
	oldUpdatedAt := user.UpdatedAt
	time.Sleep(time.Millisecond) // Ensure time difference
	user.Name = "Updated User"
	err = repo.Update(user)
	assert.NoError(t, err)

	// Verify the update
	found, err := repo.GetByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated User", found.Name)
	assert.True(t, found.UpdatedAt.After(oldUpdatedAt))
}

func TestUserRepository_Delete(t *testing.T) {
	db := database.SetupTestDB(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	repo := NewUserRepository(db)

	// Create a test user
	user := &entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}
	err := repo.Create(user)
	assert.NoError(t, err)

	// Delete the user
	err = repo.Delete(user.ID)
	assert.NoError(t, err)

	// Verify the deletion
	found, err := repo.GetByID(user.ID)
	assert.Error(t, err)
	assert.Nil(t, found)
}

func TestUserRepository_List(t *testing.T) {
	db := database.SetupTestDB(t)
	defer database.CleanupTestDB(t, db)
	defer cleanupTestData(t, db)

	repo := NewUserRepository(db)

	// Create test users
	users := []*entity.User{
		{Name: "User 1", Email: "user1@example.com", Password: "password123"},
		{Name: "User 2", Email: "user2@example.com", Password: "password123"},
		{Name: "User 3", Email: "user3@example.com", Password: "password123"},
	}

	for _, user := range users {
		err := repo.Create(user)
		assert.NoError(t, err)
	}

	// Test listing with pagination
	found, err := repo.List(1, 2)
	assert.NoError(t, err)
	assert.Len(t, found, 2)

	// Test listing all users
	found, err = repo.List(1, 10)
	assert.NoError(t, err)
	assert.Len(t, found, len(users))

	// Verify order (should be by ID ascending)
	for i := 0; i < len(found)-1; i++ {
		assert.True(t, found[i].ID < found[i+1].ID)
	}
}

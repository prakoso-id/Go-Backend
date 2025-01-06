package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	// Load .env file if it exists
	if err := godotenv.Load("../../../.env"); err != nil {
		// It's okay if .env doesn't exist in test environment
		fmt.Printf("Warning: .env file not found for tests: %v\n", err)
	}

	// Set default test environment variables if not set
	if os.Getenv("DB_HOST") == "" {
		os.Setenv("DB_HOST", "192.168.110.21")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "5432")
	}
	if os.Getenv("DB_USER") == "" {
		os.Setenv("DB_USER", "postgres-prakoso")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "@Rikudo31")
	}
	if os.Getenv("DB_NAME") == "" {
		os.Setenv("DB_NAME", "go_backend")
	}
	if os.Getenv("JWT_SECRET") == "" {
		os.Setenv("JWT_SECRET", "test_secret")
	}
}

func SetupTestDB(t *testing.T) *gorm.DB {
	dbName := os.Getenv("DB_NAME") + "_test"
	
	// Create test database if it doesn't exist
	err := createTestDatabase(dbName)
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Run migrations
	if err := Migrate(db); err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	return db
}

func createTestDatabase(dbName string) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %v", err)
	}

	// Check if database exists
	var count int64
	err = db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", dbName).Count(&count).Error
	if err != nil {
		return fmt.Errorf("failed to check database existence: %v", err)
	}

	// Create database if it doesn't exist
	if count == 0 {
		err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
		if err != nil {
			return fmt.Errorf("failed to create database: %v", err)
		}
	}

	return nil
}

func CleanupTestDB(t *testing.T, db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		t.Errorf("Failed to get underlying *sql.DB: %v", err)
		return
	}
	sqlDB.Close()
}

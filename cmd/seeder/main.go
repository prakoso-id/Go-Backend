package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-backend/internal/modules/user/seeder"
)

func initDB() (*gorm.DB, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Get database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Create database connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName, dbPort)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

func main() {
	// Parse command line flags
	seedType := flag.String("type", "all", "type of seed to run (all, admin)")
	flag.Parse()

	// Initialize database connection
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run seeders based on type
	switch *seedType {
	case "all":
		fmt.Println("Running all seeders...")
		if err := seeder.SeedAdminUsers(db); err != nil {
			log.Fatalf("Failed to seed admin users: %v", err)
		}
		// Add more seeders here as needed
		fmt.Println("All seeders completed successfully!")

	case "admin":
		fmt.Println("Running admin users seeder...")
		if err := seeder.SeedAdminUsers(db); err != nil {
			log.Fatalf("Failed to seed admin users: %v", err)
		}
		fmt.Println("Admin users seeded successfully!")

	default:
		fmt.Printf("Unknown seeder type: %s\n", *seedType)
		fmt.Println("Available types: all, admin")
		os.Exit(1)
	}
}

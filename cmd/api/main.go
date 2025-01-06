package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go-backend/internal/infrastructure/database"
	"go-backend/internal/interfaces/http/router"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Setup router
	r := router.NewRouter(db)
	r.SetupRoutes()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

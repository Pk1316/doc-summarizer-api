package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/file-processing-api/config"
	"github.com/yourusername/file-processing-api/handlers"
	"github.com/yourusername/file-processing-api/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the Gemini service with the API key
	geminiService := services.NewGeminiService(cfg.APIKey)

	// Set up the Gin router
	router := gin.Default()

	// Register routes
	handlers.RegisterRoutes(router, geminiService)

	// Start the server
	log.Printf("Server starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}

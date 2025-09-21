package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database connection
	if err := InitDB(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer CloseDB()

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	router := gin.Default()

	// Initialize routes
	setupRoutes(router)

	// Get port from environment or default to 8081
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Bus Management Service starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(router *gin.Engine) {
	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "bus-management"})
	})

	// API routes
	api := router.Group("/api")
	{
		// Bus routes
		api.POST("/buses", handleCreateBus)
		api.GET("/buses", handleGetBuses)
		api.GET("/buses/:id", handleGetBus)
		api.PUT("/buses/:id", handleUpdateBus)
		api.DELETE("/buses/:id", handleDeleteBus)

		// Staff routes
		api.POST("/staff", handleCreateStaff)
		api.GET("/staff", handleGetStaff)
		api.GET("/staff/:id", handleGetStaffMember)
		api.PUT("/staff/:id", handleUpdateStaff)
		api.DELETE("/staff/:id", handleDeleteStaff)
	}
}

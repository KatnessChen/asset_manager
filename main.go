package main

import (
	"log"

	"asset_manager/database"
	"asset_manager/middleware"
	"asset_manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	r := gin.Default()

	// Use middleware to make db available to handlers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Add custom CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Add JSON content type check middleware
	r.Use(middleware.JSONContentTypeMiddleware())
	
	// Auto Migrate the models
	// db.AutoMigrate(&models.User{}, &models.AssetType{})

	routes.SetupRoutes(r, db)

	r.Run(":8080")
}
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"asset_manager/internal/handlers"
	"asset_manager/internal/repository"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	assetTypeRepo := repository.NewAssetTypeRepository(db)
	assetTypeHandler := handlers.NewAssetTypeHandler(assetTypeRepo)

	assetRepo := repository.NewAssetRepository(db)
	assetHandler := handlers.NewAssetHandler(assetRepo)

	assetRecordRepo := repository.NewAssetRecordRepository(db)
	assetRecordHandler := handlers.NewAssetRecordHandler(assetRecordRepo)

	// User routes
	r.GET("/users/:id", userHandler.GetUserByID)
	r.GET("/users/:id/assets", assetHandler.GetAssetsByUser)

	// Asset Type routes
	r.GET("/asset-types", assetTypeHandler.GetAll)
	r.GET("/asset-types/:id", assetTypeHandler.GetByID)
	r.POST("/asset-types", assetTypeHandler.Create)
	r.PUT("/asset-types/:id", assetTypeHandler.Update)
	r.DELETE("/asset-types/:id", assetTypeHandler.Delete)

	// Asset routes
	// r.GET("/assets", assetHandler.GetAll)
	// r.GET("/assets/:id", assetHandler.GetByID)
	r.POST("/assets", assetHandler.Create)
	r.PUT("/assets/:id", assetHandler.Update)
	r.DELETE("/assets/:id", assetHandler.Delete)

	// Asset Record routes
	r.GET("/asset-records", assetRecordHandler.GetAll)
	r.GET("/asset-records/:id", assetRecordHandler.GetByID)
	r.POST("/asset-records", assetRecordHandler.Create)
	r.PUT("/asset-records/:id", assetRecordHandler.Update)
	r.DELETE("/asset-records/:id", assetRecordHandler.Delete)
}

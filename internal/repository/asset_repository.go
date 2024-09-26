// internal/repository/asset_repository.go
package repository

import (
	"asset_manager/internal/models"

	"gorm.io/gorm"
)

type AssetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *AssetRepository {
	return &AssetRepository{db: db}
}

func (r *AssetRepository) Create(asset *models.Asset) error {
	return r.db.Create(asset).Error
}

func (r *AssetRepository) GetAll() ([]models.Asset, error) {
	var assets []models.Asset
	err := r.db.Find(&assets).Error
	return assets, err
}

func (r *AssetRepository) GetByID(id uint) (models.Asset, error) {
	var asset models.Asset
	err := r.db.First(&asset, id).Error
	return asset, err
}

func (r *AssetRepository) Update(asset *models.Asset) error {
	return r.db.Save(asset).Error
}

func (r *AssetRepository) Delete(id uint) error {
	return r.db.Delete(&models.Asset{}, id).Error
}

func (r *AssetRepository) GetAssetsByUserID(userID uint) ([]models.Asset, error) {
	var assets []models.Asset
	err := r.db.Model(&models.Asset{}).
							Where(&models.Asset{UserId: userID}).
							Preload("AssetRecords").
							Find(&assets).Error
	return assets, err
}
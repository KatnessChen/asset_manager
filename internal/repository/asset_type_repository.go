package repository

import (
	"asset_manager/internal/models"

	"gorm.io/gorm"
)

type AssetTypeRepository struct {
	db *gorm.DB
}

func NewAssetTypeRepository(db *gorm.DB) *AssetTypeRepository {
	return &AssetTypeRepository{db: db}
}

func (r *AssetTypeRepository) Create(assetType *models.AssetType) error {
	return r.db.Create(assetType).Error
}

func (r *AssetTypeRepository) GetAll() ([]models.AssetType, error) {
	var assetTypes []models.AssetType
	result := r.db.Find(&assetTypes)
	return assetTypes, result.Error
}

func (r *AssetTypeRepository) GetByID(id uint) (models.AssetType, error) {
	var assetType models.AssetType
	err := r.db.First(&assetType, id).Error
	return assetType, err
}

func (r *AssetTypeRepository) Update(assetType *models.AssetType) error {
	return r.db.Save(assetType).Error
}

func (r *AssetTypeRepository) Delete(id uint) error {
	return r.db.Delete(&models.AssetType{}, id).Error
}
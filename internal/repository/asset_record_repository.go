package repository

import (
	"asset_manager/internal/models"

	"gorm.io/gorm"
)

type AssetRecordRepository struct {
	db *gorm.DB
}

func NewAssetRecordRepository(db *gorm.DB) *AssetRecordRepository {
	return &AssetRecordRepository{db: db}
}

func (r *AssetRecordRepository) Create(record *models.AssetRecord) error {
	return r.db.Create(record).Error
}

func (r *AssetRecordRepository) GetAll() ([]models.AssetRecord, error) {
	var records []models.AssetRecord
	err := r.db.Find(&records).Error
	return records, err
}

func (r *AssetRecordRepository) GetByID(id uint) (models.AssetRecord, error) {
	var record models.AssetRecord
	err := r.db.First(&record, id).Error
	return record, err
}

func (r *AssetRecordRepository) Update(record *models.AssetRecord) error {
	return r.db.Save(record).Error
}

func (r *AssetRecordRepository) Delete(id uint) error {
	return r.db.Delete(&models.AssetRecord{}, id).Error
}

func (r *AssetRecordRepository) GetByAssetID(assetID uint) ([]models.AssetRecord, error) {
	var records []models.AssetRecord
	err := r.db.Where("asset_id = ?", assetID).Find(&records).Error
	return records, err
}
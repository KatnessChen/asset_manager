package models

import (
	"time"
)

type Asset struct {
	AssetID     uint      `gorm:"primaryKey;autoIncrement" json:"asset_id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	AssetTypeID uint      `gorm:"not null" json:"asset_type_id"`
	AssetType   AssetType     `gorm:"foreignKey:AssetTypeID"`
	AssetRecords []AssetRecord `gorm:"foreignKey:AssetID"`
	Name        string    `gorm:"not null" json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
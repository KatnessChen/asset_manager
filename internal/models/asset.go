package models

import (
	"time"
)

type Asset struct {
	AssetId     uint      `gorm:"primaryKey;autoIncrement" json:"asset_id"`
	UserId      uint      `gorm:"not null" json:"user_id"`
	AssetTypeId uint      `gorm:"not null" json:"asset_type_id"`
	AssetRecords []AssetRecord `gorm:"foreignKey:AssetId" json:"asset_records"`
	Name        string    `gorm:"not null" json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
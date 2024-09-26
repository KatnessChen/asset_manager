package models

type AssetType struct {
	AssetTypeId uint      `gorm:"primaryKey;autoIncrement" json:"asset_type_id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Description string    `json:"description"`
}
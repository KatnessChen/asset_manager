package models

type AssetType struct {
	AssetTypeID uint      `gorm:"primaryKey;autoIncrement" json:"asset_type_id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Description string    `json:"description"`
}
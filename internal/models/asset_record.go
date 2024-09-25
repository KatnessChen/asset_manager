package models

import (
	"time"
)

type AssetRecord struct {
	RecordID   uint      `gorm:"column:record_id;primaryKey;autoIncrement" json:"record_id"`
	AssetID    uint      `gorm:"column:asset_id;not null" json:"asset_id"`
	Unit       float64   `gorm:"column:unit;type:decimal(15,2);not null" json:"unit"`
	UnitCost   float64   `gorm:"column:unit_cost;type:decimal(15,2);not null" json:"unit_cost"`
	UnitPrice  float64   `gorm:"column:unit_price;type:decimal(15,2);not null" json:"unit_price"`
	RecordDate time.Time `gorm:"column:record_date;type:date;not null" json:"record_date"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	// Asset      Asset     `gorm:"foreignKey:AssetID" json:"-"`
}

type LatestAssetRecord struct {
	AssetID    uint      `json:"asset_id"`
	AssetName  string    `json:"asset_name"`
	AssetTypeID  uint    `json:"asset_type_id"`
	Unit       float64   `json:"unit"`
	UnitCost   float64   `json:"unit_cost"`
	UnitPrice  float64   `json:"unit_price"`
	Value 		 float64	 `json:"value"`
	ProfitLoss float64   `json:"profit_loss"`
	RecordDate time.Time `json:"record_date"`
}
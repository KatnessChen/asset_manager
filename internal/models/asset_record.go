package models

import (
	"time"
)

type AssetRecord struct {
	RecordId   uint      `gorm:"column:record_id;primaryKey;autoIncrement" json:"record_id"`
	AssetId    uint      `gorm:"column:asset_id;not null" json:"asset_id"`
	Unit       float64   `gorm:"column:unit;type:decimal(15,2);not null" json:"unit"`
	UnitCost   float64   `gorm:"column:unit_cost;type:decimal(15,2);not null" json:"unit_cost"`
	UnitPrice  float64   `gorm:"column:unit_price;type:decimal(15,2);not null" json:"unit_price"`
	RecordDate time.Time `gorm:"column:record_date;type:date;not null" json:"record_date"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
package models

import (
	"time"
)

type User struct {
    UserId         uint         `gorm:"primaryKey" json:"user_id"`
    Username     string         `gorm:"uniqueIndex;not null" json:"user_name"`
    Assets      []Asset         `gorm:"foreignKey:UserID"`
    Email        string         `gorm:"uniqueIndex;not null" json:"email"`
    PasswordHash string         `gorm:"not null"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
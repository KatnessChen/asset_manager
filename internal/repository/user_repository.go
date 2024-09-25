package repository

import (
	"asset_manager/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
    var users []models.User
    result := r.db.Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }
    return users, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    err := r.db.Preload("Assets").
               Preload("Assets.AssetType").
               Preload("Assets.AssetRecords", func(db *gorm.DB) *gorm.DB {
                   return db.Order("record_date DESC")
               }).
               First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
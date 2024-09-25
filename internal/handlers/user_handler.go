package handlers

import (
	"net/http"
	"strconv"

	"asset_manager/internal/repository"

	"asset_manager/internal/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userRepo.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_name": user.Username,
		"user_id":   user.UserId,
		"email":     user.Email,
	})
}

func (h *UserHandler) GetLatestAssetRecords(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userRepo.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if len(user.Assets) == 0 {
		c.JSON(http.StatusOK, []models.LatestAssetRecord{})
		return
	}
	
	var latestRecords []models.LatestAssetRecord
	for _, asset := range user.Assets {
		if len(asset.AssetRecords) == 0 {
			continue
		}

		latestRecord := asset.AssetRecords[len(asset.AssetRecords)-1]
		value := latestRecord.Unit * latestRecord.UnitPrice
		profitLoss := latestRecord.Unit * (latestRecord.UnitPrice - latestRecord.UnitCost)

		latestRecords = append(latestRecords, models.LatestAssetRecord{
			AssetID:    asset.AssetID,
			AssetName:  asset.Name,
			AssetTypeID: asset.AssetType.AssetTypeID,
			Unit:       latestRecord.Unit,
			UnitCost:   latestRecord.UnitCost,
			UnitPrice:  latestRecord.UnitPrice,
			Value:      value,
			ProfitLoss: profitLoss,
			RecordDate: latestRecord.RecordDate,
		})
	}

	if len(latestRecords) == 0 {
		c.JSON(http.StatusOK, []models.LatestAssetRecord{})
		return
	}

	c.JSON(http.StatusOK, latestRecords)
}
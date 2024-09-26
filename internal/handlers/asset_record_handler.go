package handlers

import (
	"net/http"
	"strconv"
	"time"

	"asset_manager/internal/models"
	"asset_manager/internal/repository"

	"github.com/gin-gonic/gin"
)

type AssetRecordHandler struct {
	repo *repository.AssetRecordRepository
}

func NewAssetRecordHandler(repo *repository.AssetRecordRepository) *AssetRecordHandler {
	return &AssetRecordHandler{repo: repo}
}

func (h *AssetRecordHandler) Create(c *gin.Context) {
	var record models.AssetRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse and format the date
	date, err := time.Parse("2006-01-02 00:00:00", record.RecordDate.Format("2006-01-02 00:00:00"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}
	record.RecordDate = date.Truncate(24 * time.Hour)

	if err := h.repo.Create(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create asset record"})
		return
	}

	c.JSON(http.StatusCreated, record)
}

func (h *AssetRecordHandler) GetAll(c *gin.Context) {
	records, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve asset records"})
		return
	}

	c.JSON(http.StatusOK, records)
}

func (h *AssetRecordHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	record, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset record not found"})
		return
	}

	c.JSON(http.StatusOK, record)
}

func (h *AssetRecordHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var record models.AssetRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.RecordId = uint(id)
	if err := h.repo.Update(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset record"})
		return
	}

	c.JSON(http.StatusOK, record)
}

func (h *AssetRecordHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete asset record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset record deleted successfully"})
}

func (h *AssetRecordHandler) GetByAssetID(c *gin.Context) {
	assetID, err := strconv.ParseUint(c.Param("asset_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	records, err := h.repo.GetByAssetID(uint(assetID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve asset records"})
		return
	}

	c.JSON(http.StatusOK, records)
}

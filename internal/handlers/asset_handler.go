package handlers

import (
	"net/http"
	"strconv"

	"asset_manager/internal/models"
	"asset_manager/internal/repository"

	"github.com/gin-gonic/gin"
)

type AssetHandler struct {
	repo *repository.AssetRepository
}

func NewAssetHandler(repo *repository.AssetRepository) *AssetHandler {
	return &AssetHandler{repo: repo}
}

func (h *AssetHandler) Create(c *gin.Context) {
	var asset models.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&asset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create asset"})
		return
	}

	c.JSON(http.StatusCreated, asset)
}

func (h *AssetHandler) GetAll(c *gin.Context) {
	assets, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve assets"})
		return
	}

	c.JSON(http.StatusOK, assets)
}

func (h *AssetHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	asset, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	c.JSON(http.StatusOK, asset)
}

func (h *AssetHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var asset models.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	asset.AssetID = uint(id)
	if err := h.repo.Update(&asset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
		return
	}

	c.JSON(http.StatusOK, asset)
}

func (h *AssetHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete asset"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset deleted successfully"})
}

func (h *AssetHandler) GetByUserID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	assets, err := h.repo.GetByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve assets"})
		return
	}

	c.JSON(http.StatusOK, assets)
}

func (h *AssetHandler) GetAssetsByUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	assets, err := h.repo.GetByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve assets"})
		return
	}

	var simplifiedAssets []gin.H
	if len(assets) == 0 {
		c.JSON(http.StatusOK, []gin.H{})
		return
	}

	for _, asset := range assets {
		simplifiedAssets = append(simplifiedAssets, gin.H{
			"asset_id":      asset.AssetID,
			"asset_type_id": asset.AssetTypeID,
			"name":          asset.Name,
		})
	}
	c.JSON(http.StatusOK, simplifiedAssets)
}
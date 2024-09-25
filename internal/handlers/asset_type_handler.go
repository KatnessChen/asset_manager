package handlers

import (
	"net/http"
	"strconv"

	"asset_manager/internal/models"
	"asset_manager/internal/repository"

	"github.com/gin-gonic/gin"
)

type AssetTypeHandler struct {
	repo *repository.AssetTypeRepository
}

func NewAssetTypeHandler(repo *repository.AssetTypeRepository) *AssetTypeHandler {
	return &AssetTypeHandler{repo: repo}
}

func (h *AssetTypeHandler) Create(c *gin.Context) {
	var assetType models.AssetType
	if err := c.ShouldBindJSON(&assetType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&assetType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create asset type"})
		return
	}

	c.JSON(http.StatusCreated, assetType)
}

func (h *AssetTypeHandler) GetAll(c *gin.Context) {
	assetTypes, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch asset types"})
		return
	}
	c.JSON(http.StatusOK, assetTypes)
}

func (h *AssetTypeHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	assetType, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset type not found"})
		return
	}

	c.JSON(http.StatusOK, assetType)
}

func (h *AssetTypeHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var assetType models.AssetType
	if err := c.ShouldBindJSON(&assetType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assetType.AssetTypeID = uint(id)
	if err := h.repo.Update(&assetType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset type"})
		return
	}

	c.JSON(http.StatusOK, assetType)
}

func (h *AssetTypeHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete asset type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset type deleted successfully"})
}
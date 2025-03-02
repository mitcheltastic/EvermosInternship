package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/config"
	"github.com/mitcheltastic/EvermosInternship/models"
)

func GetStores(c *gin.Context) {
	var stores []models.Store
	config.DB.Find(&stores)
	c.JSON(http.StatusOK, stores)
}

func GetStoreByID(c *gin.Context) {
	storeID := c.Param("id")
	var store models.Store

	if err := config.DB.First(&store, storeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	c.JSON(http.StatusOK, store)
}

func UpdateStore(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	storeID := c.Param("id")
	var store models.Store

	// Find store and ensure user is the owner
	if err := config.DB.Where("id = ? AND user_id = ?", storeID, userID).First(&store).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Store not found or not authorized"})
		return
	}

	var input struct {
		Name     string `json:"name"`
		ImageURL string `json:"image_url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update only provided fields
	if input.Name != "" {
		store.Name = input.Name
	}
	if input.ImageURL != "" {
		store.ImageURL = input.ImageURL
	}

	config.DB.Save(&store)

	c.JSON(http.StatusOK, gin.H{"message": "Store updated successfully"})
}


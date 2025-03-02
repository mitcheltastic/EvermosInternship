package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/config"
	"github.com/mitcheltastic/EvermosInternship/models"
)

// 🔹 Get All Categories (Public)
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// 🔹 Create Category (Admin Only)
func CreateCategory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if user is an admin
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil || !user.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. Only admins can create categories."})
		return
	}

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully", "category_id": input.ID})
}

// 🔹 Update Category (Admin Only)
func UpdateCategory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if user is an admin
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil || !user.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. Only admins can update categories."})
		return
	}

	categoryID := c.Param("id")
	var category models.Category

	// Find the category
	if err := config.DB.First(&category, categoryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var input struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update category name
	category.Name = input.Name
	config.DB.Save(&category)

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// 🔹 Delete Category (Admin Only)
func DeleteCategory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if user is an admin
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil || !user.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. Only admins can delete categories."})
		return
	}

	categoryID := c.Param("id")
	if err := config.DB.Delete(&models.Category{}, categoryID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

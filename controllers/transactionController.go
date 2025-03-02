package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/config"
	"github.com/mitcheltastic/EvermosInternship/models"
	"github.com/mitcheltastic/EvermosInternship/middleware"
)

// GetTransactions retrieves all transactions for the authenticated user
func GetTransactions(c *gin.Context) {
	userID, exists := middleware.ExtractUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var transactions []models.Transaction
	config.DB.Where("user_id = ?", userID).Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

// GetTransaction retrieves a single transaction by ID
func GetTransaction(c *gin.Context) {
	userID, exists := middleware.ExtractUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	if err := config.DB.Where("id = ? AND user_id = ?", transactionID, userID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// CreateTransaction handles transaction creation
func CreateTransaction(c *gin.Context) {
	userID, exists := middleware.ExtractUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	transaction.UserID = userID
	if err := config.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// UpdateTransaction modifies an existing transaction
func UpdateTransaction(c *gin.Context) {
	userID, exists := middleware.ExtractUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	if err := config.DB.Where("id = ? AND user_id = ?", transactionID, userID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := config.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction removes a transaction
func DeleteTransaction(c *gin.Context) {
	userID, exists := middleware.ExtractUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	if err := config.DB.Where("id = ? AND user_id = ?", transactionID, userID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if err := config.DB.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

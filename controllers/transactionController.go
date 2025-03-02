package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/config"
	"github.com/mitcheltastic/EvermosInternship/models"
)

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	config.DB.Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

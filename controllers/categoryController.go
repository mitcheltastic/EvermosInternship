package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/config"
	"github.com/mitcheltastic/EvermosInternship/models"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

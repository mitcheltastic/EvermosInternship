package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/config"
	"github.com/mitcheltastic/EvermosInternship/models"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

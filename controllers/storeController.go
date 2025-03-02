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

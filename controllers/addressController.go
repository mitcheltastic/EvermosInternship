package controllers

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/config"
	"github.com/mitcheltastic/EvermosInternship/models"
)

// Province struct to store API response
type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// City struct to store API response
type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// 🔹 Get Provinces from EMSIFA API
func GetProvinces(c *gin.Context) {
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch provinces"})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var provinces []Province
	json.Unmarshal(body, &provinces)

	c.JSON(http.StatusOK, gin.H{"provinces": provinces})
}

// 🔹 Get Cities by Province ID from EMSIFA API
func GetCitiesByProvince(c *gin.Context) {
	provinceID := c.Param("province_id")
	url := "https://www.emsifa.com/api-wilayah-indonesia/api/regencies/" + provinceID + ".json"

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cities"})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var cities []City
	json.Unmarshal(body, &cities)

	c.JSON(http.StatusOK, gin.H{"cities": cities})
}

// 🔹 Get All User Addresses
func GetUserAddresses(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	fmt.Println("🔍 DEBUG: Fetching addresses for user ID:", userID) // 🔥 Debugging Line

	var addresses []models.Address
	err := config.DB.Where("user_id = ?", userID).Find(&addresses).Error
	if err != nil {
		fmt.Println("❌ ERROR:", err) // 🔥 Debugging Line
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch addresses", "details": err.Error()})
		return
	}

	fmt.Println("✅ Addresses Found:", addresses) // 🔥 Debugging Line
	c.JSON(http.StatusOK, gin.H{"addresses": addresses})
}

// 🔹 Add New Address
func AddAddress(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input models.Address
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = userID.(uint)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address added successfully", "address_id": input.ID})
}

// 🔹 Update Address
func UpdateAddress(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	addressID := c.Param("id")
	var address models.Address

	// Ensure the user owns the address
	if err := config.DB.Where("id = ? AND user_id = ?", addressID, userID).First(&address).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Address not found or not authorized"})
		return
	}

	var input struct {
		Detail     string `json:"detail"`
		PostalCode string `json:"postal_code"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided
	if input.Detail != "" {
		address.Detail = input.Detail
	}
	if input.PostalCode != "" {
		address.PostalCode = input.PostalCode
	}

	config.DB.Save(&address)

	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}

// 🔹 Delete Address
func DeleteAddress(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	addressID := c.Param("id")
	if err := config.DB.Where("id = ? AND user_id = ?", addressID, userID).Delete(&models.Address{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
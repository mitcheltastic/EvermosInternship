package controllers

import (
	"fmt" // ✅ Added fmt for debugging
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/config"
	"github.com/mitcheltastic/EvermosInternship/models"
	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 🔥 Debug: Print received data
	fmt.Println("Received JSON Data:", input)

	// 🔥 Ensure password is not empty
	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	// 🔥 Hash password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	fmt.Println("🔹 Storing Hashed Password:", string(hashedPassword)) // Debugging

	input.Password = string(hashedPassword)

	// 🔥 Store user in the database
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	// 🚀 Auto-create store
	store := models.Store{
		UserID:   input.ID,
		Name:     input.Name + "'s Store",
		ImageURL: "default-store-image.png",
	}
	if err := config.DB.Create(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User registered, but failed to create store", "details": err.Error()})
		return
	}

	// Generate JWT token
	token, err := config.GenerateToken(input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// ✅ Success response
	c.JSON(http.StatusOK, gin.H{"token": token, "store_id": store.ID})
}

// Login handles user authentication
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email (make sure email comparison is case-insensitive)
	var user models.User
	if err := config.DB.Where("LOWER(email) = LOWER(?)", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := config.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"token": token})
}


// Logout endpoint
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out. Please remove the token from your client storage."})
}



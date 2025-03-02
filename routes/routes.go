package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/controllers"
	"github.com/mitcheltastic/EvermosInternship/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Authentication routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	// Public Routes (No authentication needed)
	r.GET("/addresses/provinces", controllers.GetProvinces) // ✅ Get Provinces
	r.GET("/addresses/cities/:province_id", controllers.GetCitiesByProvince) // ✅ Get Cities
	// Public Routes
	r.GET("/categories", controllers.GetCategories) // ✅ Get Categories (Public)

	// Protected Routes (Require Authentication)
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		// User routes
		auth.GET("/users/profile", controllers.GetUserProfile)       // Get user profile
		auth.PUT("/users/profile", controllers.UpdateUserProfile)    // Update profile
		auth.PUT("/users/change-password", controllers.ChangePassword) // Change password
		auth.DELETE("/users/delete", controllers.DeleteAccount)      // Delete account

		// Store routes
		auth.GET("/stores", controllers.GetStores) // ✅ Get All Stores
		auth.GET("/stores/:id", controllers.GetStoreByID) // ✅ Get Store by ID
		auth.PUT("/stores/:id", controllers.UpdateStore) // ✅ Update Store

		// Address routes
		auth.GET("/addresses", controllers.GetUserAddresses) // ✅ Get All Addresses
		auth.POST("/addresses", controllers.AddAddress) // ✅ Add Address
		auth.PUT("/addresses/:id", controllers.UpdateAddress) // ✅ Update Address
		auth.DELETE("/addresses/:id", controllers.DeleteAddress) // ✅ Delete Address

		// Admin-Only Routes
		auth.POST("/categories", controllers.CreateCategory) // ✅ Create Category (Admin)
		auth.PUT("/categories/:id", controllers.UpdateCategory) // ✅ Update Category (Admin)
		auth.DELETE("/categories/:id", controllers.DeleteCategory) // ✅ Delete Category (Admin)

		// Product routes
		auth.GET("/products", controllers.GetProducts)

		// Transaction routes
		auth.GET("/transactions", controllers.GetTransactions)
	}

	return r
}

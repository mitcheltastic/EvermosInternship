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

	// Protected Routes (Require Authentication)
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// User routes
		auth.GET("/users/profile", controllers.GetUserProfile)       // Get user profile
		auth.PUT("/users/profile", controllers.UpdateUserProfile)    // Update profile
		auth.PUT("/users/change-password", controllers.ChangePassword) // Change password
		auth.DELETE("/users/delete", controllers.DeleteAccount)      // Delete account

		// Store routes
		auth.GET("/stores", controllers.GetStores)

		// Category routes
		auth.GET("/categories", controllers.GetCategories)

		// Product routes
		auth.GET("/products", controllers.GetProducts)

		// Transaction routes
		auth.GET("/transactions", controllers.GetTransactions)
	}

	return r
}

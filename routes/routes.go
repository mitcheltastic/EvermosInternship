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
	r.GET("/addresses/provinces", controllers.GetProvinces) 
	r.GET("/addresses/cities/:province_id", controllers.GetCitiesByProvince) 
	r.GET("/categories", controllers.GetCategories) 
	r.GET("/products", controllers.GetProducts) 
	r.GET("/products/:id", controllers.GetProductByID) 

	// Protected Routes (Require Authentication)
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		// User routes
		auth.GET("/users/profile", controllers.GetUserProfile)       
		auth.PUT("/users/profile", controllers.UpdateUserProfile)    
		auth.PUT("/users/change-password", controllers.ChangePassword) 
		auth.DELETE("/users/delete", controllers.DeleteAccount)      

		// Store routes
		auth.GET("/stores", controllers.GetStores) 
		auth.GET("/stores/:id", controllers.GetStoreByID) 
		auth.PUT("/stores/:id", controllers.UpdateStore) 

		// Address routes
		auth.GET("/addresses", controllers.GetUserAddresses) 
		auth.POST("/addresses", controllers.AddAddress) 
		auth.PUT("/addresses/:id", controllers.UpdateAddress) 
		auth.DELETE("/addresses/:id", controllers.DeleteAddress) 

		// Admin-Only Routes
		auth.POST("/categories", controllers.CreateCategory) 
		auth.PUT("/categories/:id", controllers.UpdateCategory) 
		auth.DELETE("/categories/:id", controllers.DeleteCategory) 

		// Store Owner Routes
		auth.POST("/products", controllers.CreateProduct) 
		auth.PUT("/products/:id", controllers.UpdateProduct) 
		auth.DELETE("/products/:id", controllers.DeleteProduct) 

		// Transaction routes (Added Here)
		auth.GET("/transactions", controllers.GetTransactions) // Get all transactions of the logged-in user
		auth.GET("/transactions/:id", controllers.GetTransaction) // Get a specific transaction by ID
		auth.POST("/transactions", controllers.CreateTransaction) // Create a new transaction
		auth.PUT("/transactions/:id", controllers.UpdateTransaction) // Update an existing transaction
		auth.DELETE("/transactions/:id", controllers.DeleteTransaction) // Delete a transaction
	}

	return r
}

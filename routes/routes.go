package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mitcheltastic/EvermosInternship/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Authentication
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	return r
}

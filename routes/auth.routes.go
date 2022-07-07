package routes

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) *gin.RouterGroup {
	authController := controllers.NewAuthController()
	auth := incomingRoutes.Group("/auth")
	{
		auth.POST("/register", authController.Register())
		auth.POST("/login", authController.Login())
	}
	return auth
}

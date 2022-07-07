package routes

import (
	"myapp/controllers"
	"myapp/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userController := controllers.NewUserController()
	var user *gin.RouterGroup = incomingRoutes.Group("/user")
	{
		user.GET("/", userController.FindAll())
		user.GET("/:id", userController.FindOne())
		user.PUT("/:id", middlewares.AuthorizeJWT(), userController.UpdateOne())
		user.DELETE("/:id", middlewares.AuthorizeJWT(), userController.DeleteOne())
	}
}

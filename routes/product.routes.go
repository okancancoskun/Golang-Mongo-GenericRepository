package routes

import (
	"myapp/controllers"
	"myapp/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine) *gin.RouterGroup {
	productController := controllers.NewProductController()
	var product *gin.RouterGroup = incomingRoutes.Group("/product")
	{
		product.GET("/", productController.FindAll())
		product.GET("/:id", productController.FindOne())
		product.POST("/create", middlewares.AuthorizeJWT(), productController.Create())
		product.DELETE("/:id", middlewares.AuthorizeJWT(), productController.DeleteOne())
		product.PUT("/:id", middlewares.AuthorizeJWT(), productController.UpdateOne())
	}
	return product
}

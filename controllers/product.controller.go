package controllers

import (
	"myapp/models"
	"myapp/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{productService: *services.NewProductService()}
}

func (controller *ProductController) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		prds, err := controller.productService.FindAll(bson.M{}, &options.FindOptions{})
		if err != nil {
			c.IndentedJSON(500, "Someting Went Wrong")
		}
		c.IndentedJSON(200, prds)
	}
}

func (controller *ProductController) FindOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, _ := primitive.ObjectIDFromHex(c.Param("id"))
		result, err := controller.productService.FindOne(bson.M{"_id": productId}, &options.FindOneOptions{})
		if err != nil {
			c.IndentedJSON(404, "Not Found")
		}
		c.IndentedJSON(200, result)
	}
}

func (controller *ProductController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		currentUserData, _ := c.Get("user")
		var userId = currentUserData.(map[string]interface{})["_id"].(string)
		product.ID = primitive.NewObjectID()
		product.UserId, _ = primitive.ObjectIDFromHex(userId)
		c.BindJSON(&product)
		result, err := controller.productService.Create(product, &options.InsertOneOptions{})
		if err != nil {
			c.IndentedJSON(500, "Something went wrong")
		}
		c.IndentedJSON(201, result)
	}
}

func (controller *ProductController) DeleteOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, _ := primitive.ObjectIDFromHex(c.Param("id"))
		result, err := controller.productService.DeleteOne(bson.M{"_id": productId})
		if err != nil {
			c.IndentedJSON(500, err)
		}
		c.IndentedJSON(204, result)
	}
}

func (controller *ProductController) UpdateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, _ := primitive.ObjectIDFromHex(c.Param("id"))
		var body interface{}
		c.BindJSON(&body)
		result := controller.productService.UpdateOne(bson.M{"_id": productId}, body, &options.FindOneAndUpdateOptions{})
		c.IndentedJSON(204, result)
	}
}

package controllers

import (
	"myapp/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserController struct {
	userService services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: *services.NewUserService(),
	}
}

func (controller *UserController) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := controller.userService.FindAll(bson.M{}, &options.FindOptions{})
		if err != nil {
			c.IndentedJSON(500, "Someting Went Wrong")
		}
		c.IndentedJSON(200, users)
	}
}

func (controller *UserController) FindOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := primitive.ObjectIDFromHex(c.Param("id"))
		result, err := controller.userService.FindOne(bson.M{"_id": userId}, &options.FindOneOptions{})
		if err != nil {
			c.IndentedJSON(404, "Not Found")
		}
		c.IndentedJSON(200, result)
	}
}

func (controller *UserController) DeleteOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, _ := primitive.ObjectIDFromHex(c.Param("id"))
		result, err := controller.userService.DeleteOne(bson.M{"_id": productId})
		if err != nil {
			c.IndentedJSON(500, err)
		}
		c.IndentedJSON(204, result)
	}
}

func (controller *UserController) UpdateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, _ := primitive.ObjectIDFromHex(c.Param("id"))
		var body interface{}
		c.BindJSON(&body)
		result := controller.userService.UpdateOne(bson.M{"_id": productId}, body, &options.FindOneAndUpdateOptions{})
		c.IndentedJSON(204, result)
	}
}

package controllers

import (
	"myapp/models"
	"myapp/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: *services.NewAuthService(),
		userService: *services.NewUserService(),
	}
}

func (controller *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		c.BindJSON(&user)
		result, err := controller.authService.Register(user)
		if err != nil {
			c.IndentedJSON(500, "Something went wrong")
		}
		c.IndentedJSON(201, result)
	}
}

func (controller *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginDto models.LoginUser
		c.BindJSON(&loginDto)
		user, err := controller.userService.FindOne(bson.M{"email": loginDto.Email}, &options.FindOneOptions{})
		if err != nil {
			c.IndentedJSON(404, "User Doesnt Exist")
		} else {
			var token string = controller.authService.GenerateJwt(user.ID, user.Email)
			c.IndentedJSON(201, map[string]interface{}{
				"_id":   user.ID,
				"email": user.Email,
				"name":  user.Name,
				"token": token,
			})
		}
	}
}

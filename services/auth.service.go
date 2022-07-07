package services

import (
	"fmt"
	"myapp/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService UserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		userService: *NewUserService(),
	}
}

func (service *AuthService) Register(body models.User) (*mongo.InsertOneResult, error) {
	_, err := service.userService.FindOne(bson.M{"email": body.Email}, &options.FindOneOptions{})
	if err != nil {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
		body.Password = string(hashed)
		body.ID = primitive.NewObjectID()
		return service.userService.Create(body, &options.InsertOneOptions{})
	} else {
		return nil, err
	}
}

func (service *AuthService) GenerateJwt(userId primitive.ObjectID, email string) string {
	claims := jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
		"iat":   time.Now().Unix(),
		"_id":   string(userId.Hex()),
		"email": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("secretjwt"))
	return tokenString
}

func ValidateToken(token string) (*jwt.Token, error) {

	//2nd arg function return secret key after checking if the signing method is HMAC and returned key is used by 'Parse' to decode the token)
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secretjwt"), nil
	})
	return parsed, err
}

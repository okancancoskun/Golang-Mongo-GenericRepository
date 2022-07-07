package services

import (
	"myapp/models"
	"myapp/repositories"
)

type UserService struct {
	*GenericService[models.User]
}

func NewUserService() *UserService {
	return &UserService{
		&GenericService[models.User]{
			repository: *repositories.NewUserRepository().GenericRepository,
		},
	}
}

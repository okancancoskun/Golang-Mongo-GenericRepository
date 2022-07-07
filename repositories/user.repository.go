package repositories

import (
	"myapp/database"
	"myapp/models"
)

type IUserRepository interface {
	*IGenericRepository[models.User]
}
type UserRepository struct {
	*GenericRepository[models.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		GenericRepository: &GenericRepository[models.User]{db: database.CollectionData(database.DBSet(), "Users")},
	}
}

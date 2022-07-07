package services

import (
	"myapp/models"
	"myapp/repositories"
)

type ProductService struct {
	*GenericService[models.Product]
}

func NewProductService() *ProductService {
	return &ProductService{
		&GenericService[models.Product]{
			repository: *repositories.NewProductRepository().GenericRepository,
		},
	}
}

// type IProductService interface {
// 	Create(body models.Product, options *options.InsertOneOptions) (*mongo.InsertOneResult, error)
// 	FindOne(filter bson.M, options *options.FindOneOptions) (models.Product, error)
// 	FindAll(filter bson.M, options *options.FindOptions) ([]models.Product, error)
// 	DeleteOne(filter bson.M) (*mongo.DeleteResult, error)
// }

// type ProductService struct {
// 	repository repositories.IProductRepository
// }

// func Init() IProductService {
// 	return &ProductService{repository: repositories.Init()}
// }

// func (service *ProductService) Create(body models.Product, options *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
// 	return service.repository.Create(body, options)
// }

// func (service *ProductService) FindOne(filter bson.M, options *options.FindOneOptions) (models.Product, error) {
// 	return service.repository.FindOne(filter, options)
// }

// func (service *ProductService) FindAll(filter bson.M, options *options.FindOptions) ([]models.Product, error) {
// 	return service.repository.FindAll(filter, options)
// }

// func (service *ProductService) DeleteOne(filter bson.M) (*mongo.DeleteResult, error) {
// 	return service.repository.DeleteOne(filter)
// }

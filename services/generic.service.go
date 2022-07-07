package services

import (
	"myapp/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GenericService[T any] struct {
	repository repositories.GenericRepository[T]
}

func (service *GenericService[T]) Create(body T, options *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return service.repository.Create(body, options)
}

func (service *GenericService[T]) FindOne(filter bson.M, options *options.FindOneOptions) (T, error) {
	return service.repository.FindOne(filter, options)
}

func (service *GenericService[T]) FindAll(filter bson.M, options *options.FindOptions) ([]T, error) {
	return service.repository.FindAll(filter, options)
}

func (service *GenericService[T]) DeleteOne(filter bson.M) (*mongo.DeleteResult, error) {
	return service.repository.DeleteOne(filter)
}

func (service *GenericService[T]) UpdateOne(filter bson.M, update interface{}, options *options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return service.repository.UpdateOne(filter, update, options)
}

package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IGenericRepository[T any] interface {
	Create(body T, options *options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(filter bson.M, options *options.FindOneOptions) (T, error)
	FindAll(filter bson.M, options *options.FindOptions) ([]T, error)
	DeleteOne(filter bson.M) (*mongo.DeleteResult, error)
}

type GenericRepository[T any] struct {
	db *mongo.Collection
}

func (repository *GenericRepository[T]) Create(body T, options *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	var ctx context.Context = context.TODO()
	return repository.db.InsertOne(ctx, body, options)
}

func (repository *GenericRepository[T]) FindOne(filter bson.M, options *options.FindOneOptions) (T, error) {
	var ctx context.Context = context.TODO()
	var data T
	cur := repository.db.FindOne(ctx, filter, options)
	err := cur.Decode(&data)
	return data, err
}

func (repository *GenericRepository[T]) FindAll(filter bson.M, options *options.FindOptions) ([]T, error) {
	var ctx context.Context = context.TODO()
	var datas []T
	cur, err := repository.db.Find(ctx, filter, options)
	/* for cur.Next(ctx) {
		var result models.Product
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, result)
	} */
	cur.All(ctx, &datas)
	return datas, err
}

func (repository *GenericRepository[T]) UpdateOne(filter bson.M, update interface{}, options *options.FindOneAndUpdateOptions) *mongo.SingleResult {
	ctx := context.TODO()
	return repository.db.FindOneAndUpdate(ctx, filter, update, options)
}

func (repository *GenericRepository[T]) DeleteOne(filter bson.M) (*mongo.DeleteResult, error) {
	ctx := context.TODO()
	return repository.db.DeleteOne(ctx, filter)
}

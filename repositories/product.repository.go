package repositories

import (
	"myapp/database"
	"myapp/models"
)

type ProductRepository struct {
	*GenericRepository[models.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		&GenericRepository[models.Product]{db: database.CollectionData(database.DBSet(), "Products")},
	}
}

// type IProductRepository interface {
// 	Create(body models.Product, options *options.InsertOneOptions) (*mongo.InsertOneResult, error)
// 	FindOne(filter bson.M, options *options.FindOneOptions) (models.Product, error)
// 	FindAll(filter bson.M, options *options.FindOptions) ([]models.Product, error)
// 	DeleteOne(filter bson.M) (*mongo.DeleteResult, error)
// }

// type ProductRepository struct {
// 	db *mongo.Collection
// }

// func Init() IProductRepository {
// 	return &ProductRepository{db: database.CollectionData(database.DBSet(), "Products")}
// }

// func (repository *ProductRepository) Create(body models.Product, options *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
// 	var ctx context.Context = context.TODO()
// 	body.ID = primitive.NewObjectID()
// 	return repository.db.InsertOne(ctx, body, options)
// }

// func (repository *ProductRepository) FindOne(filter bson.M, options *options.FindOneOptions) (models.Product, error) {
// 	var ctx context.Context = context.TODO()
// 	var product models.Product
// 	cur := repository.db.FindOne(ctx, filter, options)
// 	err := cur.Decode(&product)
// 	return product, err
// }

// func (repository *ProductRepository) FindAll(filter bson.M, options *options.FindOptions) ([]models.Product, error) {
// 	var ctx context.Context = context.TODO()
// 	var products []models.Product
// 	cur, err := repository.db.Find(ctx, filter, options)
// 	/* for cur.Next(ctx) {
// 		var result models.Product
// 		err := cur.Decode(&result)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		products = append(products, result)
// 	} */
// 	cur.All(ctx, &products)
// 	return products, err
// }

// func (repository *ProductRepository) DeleteOne(filter bson.M) (*mongo.DeleteResult, error) {
// 	ctx := context.TODO()
// 	return repository.db.DeleteOne(ctx, filter)
// }

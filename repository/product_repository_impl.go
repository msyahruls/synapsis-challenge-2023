package repository

import (
	"synapsis-challange/config"
	"synapsis-challange/helper"
	"synapsis-challange/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImpl struct {
	Collection *mongo.Collection
}

type ProductRepository interface {
	Create(product domain.Product)
	FindById(productId string) (domain.Product, error)
	FindAll(name string, category string) []domain.Product
	Update(product domain.Product, productId string)
	Delete(productId string) error
}

func NewProductRepository(database *mongo.Database) ProductRepository {
	return &ProductRepositoryImpl{
		Collection: database.Collection("products"),
	}
}

func (repository *ProductRepositoryImpl) Create(product domain.Product) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":         product.ID,
		"category_id": product.CategoryID,
		"name":        product.Name,
		"qty":         product.Qty,
		"price":       product.Price,
	})

	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(productId string) (domain.Product, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	var product domain.Product
	query := bson.M{"_id": productId}

	result := repository.Collection.FindOne(ctx, query)
	result.Decode(&product)

	return product, result.Err()
}

func (repository *ProductRepositoryImpl) FindAll(name string, category string) []domain.Product {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := helper.QueryProduct("", name, category)

	cursor, err := repository.Collection.Find(ctx, query)
	helper.PanicIfError(err)

	var products []domain.Product

	for cursor.Next(ctx) {
		var product domain.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	return products
}

func (repository *ProductRepositoryImpl) FindIsExist(name string) bool {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := helper.QueryProductExist(name)

	cursor, err := repository.Collection.Find(ctx, query)
	helper.PanicIfError(err)

	for cursor.Next(ctx) {
		return true
	}

	return false
}

func (repository *ProductRepositoryImpl) Update(product domain.Product, productId string) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := bson.M{"_id": productId}
	update := bson.M{
		"$set": bson.M{
			"category_id": product.CategoryID,
			"name":        product.Name,
			"qty":         product.Qty,
			"price":       product.Price,
		},
	}

	_, err := repository.Collection.UpdateOne(ctx, query, update)

	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) AssignPermission(product domain.Product, productId string) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := bson.M{"_id": productId}
	update := bson.M{
		"$set": bson.M{
			// "name":       product.Label,
			// "value":       product.Value,
			// "permissions": product.Permission,
		},
	}

	_, err := repository.Collection.UpdateOne(ctx, query, update)

	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) Delete(productId string) error {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.DeleteOne(ctx, bson.M{
		"_id": productId,
	})

	helper.PanicIfError(err)

	return nil
}

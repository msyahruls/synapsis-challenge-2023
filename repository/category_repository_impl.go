package repository

import (
	"synapsis-challange/config"
	"synapsis-challange/helper"
	"synapsis-challange/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepositoryImpl struct {
	Collection *mongo.Collection
}

type CategoryRepository interface {
	Create(category domain.Category)
	FindById(categoryId string) (domain.Category, error)
	FindAll(name string) []domain.Category
	FindIsExist(id string, name string) bool
	Update(category domain.Category, categoryId string)
	Delete(categoryId string) error
}

func NewCategoryRepository(database *mongo.Database) CategoryRepository {
	return &CategoryRepositoryImpl{
		Collection: database.Collection("categories"),
	}
}

func (repository *CategoryRepositoryImpl) Create(category domain.Category) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":  category.ID,
		"name": category.Name,
	})

	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(categoryId string) (domain.Category, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	var category domain.Category
	query := bson.M{"_id": categoryId}

	result := repository.Collection.FindOne(ctx, query)
	result.Decode(&category)

	return category, result.Err()
}

func (repository *CategoryRepositoryImpl) FindAll(name string) []domain.Category {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := helper.QueryCategory("", name)

	cursor, err := repository.Collection.Find(ctx, query)
	helper.PanicIfError(err)

	var categories []domain.Category

	for cursor.Next(ctx) {
		var category domain.Category
		cursor.Decode(&category)
		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) FindIsExist(id string, name string) bool {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := helper.QueryCategoryExist(id, name)

	cursor, err := repository.Collection.Find(ctx, query)
	helper.PanicIfError(err)

	for cursor.Next(ctx) {
		return true
	}

	return false
}

func (repository *CategoryRepositoryImpl) Update(category domain.Category, categoryId string) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := bson.M{"_id": categoryId}
	update := bson.M{
		"$set": bson.M{
			"name": category.Name,
		},
	}

	_, err := repository.Collection.UpdateOne(ctx, query, update)

	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) Delete(categoryId string) error {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.DeleteOne(ctx, bson.M{
		"_id": categoryId,
	})

	helper.PanicIfError(err)

	return nil
}

package repository

import (
	"synapsis-challange/config"
	"synapsis-challange/helper"
	"synapsis-challange/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepositoryImpl struct {
	Collection *mongo.Collection
}

type CartRepository interface {
	Create(cart domain.Cart)
	FindById(cartId string) (domain.Cart, error)
	FindByUserId(userId string) (domain.Cart, error)
	FindAll(UserID string, ProductID string) []domain.Cart
	FindIsExist(UserID string, ProductID string) bool
	Update(cart domain.Cart, cartId string)
	Delete(cartId string) error
	DeleteByUserId(userId string) error
}

func NewCartRepository(database *mongo.Database) CartRepository {
	return &CartRepositoryImpl{
		Collection: database.Collection("carts"),
	}
}

func (repository *CartRepositoryImpl) Create(cart domain.Cart) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":        cart.ID,
		"user_id":    cart.UserID,
		"product_id": cart.ProductID,
		"qty":        cart.Qty,
		"total":      cart.Total,
	})

	helper.PanicIfError(err)
}

func (repository *CartRepositoryImpl) FindById(cartId string) (domain.Cart, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	var cart domain.Cart
	query := bson.M{"_id": cartId}

	result := repository.Collection.FindOne(ctx, query)
	result.Decode(&cart)

	return cart, result.Err()
}

func (repository *CartRepositoryImpl) FindByUserId(userId string) (domain.Cart, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	var cart domain.Cart
	query := bson.M{"user_id": userId}

	result := repository.Collection.FindOne(ctx, query)
	result.Decode(&cart)

	return cart, result.Err()
}

func (repository *CartRepositoryImpl) FindAll(UserID string, ProductID string) []domain.Cart {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := helper.QueryCart(UserID, ProductID)

	cursor, err := repository.Collection.Find(ctx, query)
	helper.PanicIfError(err)

	var carts []domain.Cart

	for cursor.Next(ctx) {
		var cart domain.Cart
		cursor.Decode(&cart)
		carts = append(carts, cart)
	}

	return carts
}

func (repository *CartRepositoryImpl) FindIsExist(UserID string, ProductID string) bool {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := helper.QueryCartExist(UserID, ProductID)

	cursor, err := repository.Collection.Find(ctx, query)
	helper.PanicIfError(err)

	for cursor.Next(ctx) {
		return true
	}

	return false
}

func (repository *CartRepositoryImpl) Update(cart domain.Cart, cartId string) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := bson.M{"_id": cartId}
	update := bson.M{
		"$set": bson.M{
			"user_id":    cart.UserID,
			"product_id": cart.ProductID,
			"qty":        cart.Qty,
			"total":      cart.Total,
		},
	}

	_, err := repository.Collection.UpdateOne(ctx, query, update)

	helper.PanicIfError(err)
}

func (repository *CartRepositoryImpl) Delete(cartId string) error {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.DeleteOne(ctx, bson.M{
		"_id": cartId,
	})

	helper.PanicIfError(err)

	return nil
}

func (repository *CartRepositoryImpl) DeleteByUserId(userId string) error {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{
		"user_id": userId,
	})

	helper.PanicIfError(err)

	return nil
}

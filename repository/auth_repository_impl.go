package repository

import (
	"synapsis-challange/config"
	"synapsis-challange/helper"
	"synapsis-challange/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositoryImpl struct {
	Collection           *mongo.Collection
	CollectionToken      *mongo.Collection
	CollectionRespondent *mongo.Collection
}

type AuthRepository interface {
	LoginEmail(email string) (domain.User, error)
	Register(user domain.User)
	FindByQuery(query, value string) (domain.User, error)

	CreateToken(tokens domain.ResetPasswordToken)
	UpdatePassword(user domain.User)
	CheckToken(token string) (domain.ResetPasswordToken, error)
	DeleteToken(token string) error

	CreateRespondent(user domain.User)
	FindRespondentByQuery(query bson.M) (domain.User, error)
}

func NewAuthRepository(database *mongo.Database) AuthRepository {
	return &AuthRepositoryImpl{
		Collection:           database.Collection("users"),
		CollectionToken:      database.Collection("reset_token"),
		CollectionRespondent: database.Collection("respondent_user"),
	}
}

func (repository *AuthRepositoryImpl) Register(user domain.User) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":        user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"password":   user.Password,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})

	helper.PanicIfError(err)
}

func (repository *AuthRepositoryImpl) LoginEmail(email string) (domain.User, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "email", Value: email}}}}
	result, err := repository.Collection.Aggregate(ctx, mongo.Pipeline{matchStage})
	helper.PanicIfError(err)

	var user domain.User
	result.Next(ctx)
	result.Decode(&user)

	return user, result.Err()
}

func (repository *AuthRepositoryImpl) FindByQuery(query, value string) (domain.User, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	var user domain.User
	key := query
	find := bson.M{key: value}

	result := repository.Collection.FindOne(ctx, find)
	result.Decode(&user)

	return user, result.Err()
}

func (repository *AuthRepositoryImpl) CreateToken(tokens domain.ResetPasswordToken) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.CollectionToken.InsertOne(ctx, bson.M{
		"_id":        tokens.Token,
		"email":      tokens.Email,
		"phone":      tokens.Phone,
		"created_at": tokens.CreatedAt,
	})

	helper.PanicIfError(err)
}

func (repository *AuthRepositoryImpl) UpdatePassword(user domain.User) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := bson.M{"email": user.Email}
	update := bson.M{
		"$set": bson.M{
			"password": user.Password,
		},
	}
	_, err := repository.Collection.UpdateOne(ctx, query, update)

	helper.PanicIfError(err)
}

func (repository *AuthRepositoryImpl) CheckToken(token string) (domain.ResetPasswordToken, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()
	var tokens domain.ResetPasswordToken

	result := repository.CollectionToken.FindOne(ctx, bson.M{"_id": token})
	result.Decode(&tokens)

	return tokens, result.Err()
}

func (repository *AuthRepositoryImpl) DeleteToken(token string) error {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.CollectionToken.DeleteOne(ctx, bson.M{"_id": token})

	helper.PanicIfError(err)

	return nil
}

func (repository *AuthRepositoryImpl) CreateRespondent(user domain.User) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.CollectionRespondent.InsertOne(ctx, bson.M{
		"_id":   user.ID,
		"name":  user.Name,
		"email": user.Email,
		// "phone":      user.Phone,
		// "gender":     user.Gender,
		"created_at": user.CreatedAt,
	})

	helper.PanicIfError(err)
}

func (repository *AuthRepositoryImpl) FindRespondentByQuery(query bson.M) (domain.User, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	var user domain.User
	find := query

	result := repository.CollectionRespondent.FindOne(ctx, find)
	result.Decode(&user)

	return user, nil
}

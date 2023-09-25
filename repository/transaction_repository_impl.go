package repository

import (
	"synapsis-challange/config"
	"synapsis-challange/helper"
	"synapsis-challange/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepositoryImpl struct {
	Collection *mongo.Collection
}

type TransactionRepository interface {
	Create(transaction domain.Transaction)
	FindById(transactionId string) (domain.Transaction, error)
	FindAll(UserID string) []domain.Transaction
	FindIsExist(UserID string) bool
	// Update(transaction domain.Transaction, transactionId string)
	// Delete(transactionId string) error
}

func NewTransactionRepository(database *mongo.Database) TransactionRepository {
	return &TransactionRepositoryImpl{
		Collection: database.Collection("transactions"),
	}
}

func (repository *TransactionRepositoryImpl) Create(transaction domain.Transaction) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":          transaction.ID,
		"user_id":      transaction.UserID,
		"sum_total":    transaction.SumTotal,
		"payment":      transaction.Payment,
		"payback":      transaction.Payback,
		"transactions": transaction.TransactionDetails,
	})

	helper.PanicIfError(err)
}

func (repository *TransactionRepositoryImpl) FindById(transactionId string) (domain.Transaction, error) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	var transaction domain.Transaction
	query := bson.M{"_id": transactionId}

	result := repository.Collection.FindOne(ctx, query)
	result.Decode(&transaction)

	return transaction, result.Err()
}

func (repository *TransactionRepositoryImpl) FindAll(UserID string) []domain.Transaction {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := helper.QueryTransaction(UserID)

	cursor, err := repository.Collection.Find(ctx, query)
	helper.PanicIfError(err)

	var transactions []domain.Transaction

	for cursor.Next(ctx) {
		var transaction domain.Transaction
		cursor.Decode(&transaction)
		transactions = append(transactions, transaction)
	}

	return transactions
}

func (repository *TransactionRepositoryImpl) FindIsExist(UserID string) bool {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	query := helper.QueryTransactionExist(UserID)

	cursor, err := repository.Collection.Find(ctx, query)
	helper.PanicIfError(err)

	for cursor.Next(ctx) {
		return true
	}

	return false
}

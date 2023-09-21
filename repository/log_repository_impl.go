package repository

import (
	"synapsis-challange/config"
	"synapsis-challange/helper"
	"synapsis-challange/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepositoryImpl struct {
	Collection *mongo.Collection
}

type LogRepository interface {
	Create(log domain.Logs)
	FindAll(page, limit string) []domain.Logs
}

func NewLogRepository(database *mongo.Database) LogRepository {
	return &LogRepositoryImpl{
		Collection: database.Collection("logs"),
	}
}

func (repository *LogRepositoryImpl) Create(log domain.Logs) {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"actor":     log.Actor,
		"action":    log.Action,
		"timestamp": log.Timestamp,
	})

	helper.PanicIfError(err)
}

func (repository *LogRepositoryImpl) FindAll(page, limit string) []domain.Logs {
	ctx, cancel := config.NewDBContext()
	defer cancel()

	// (page, limit, sort)
	filter := helper.Pagging(page, limit, -1)

	cursor, err := repository.Collection.Find(ctx, bson.M{}, filter)
	helper.PanicIfError(err)

	var systemLogs []domain.Logs

	for cursor.Next(ctx) {
		var systemLog domain.Logs
		cursor.Decode(&systemLog)
		systemLogs = append(systemLogs, systemLog)
	}

	return systemLogs
}

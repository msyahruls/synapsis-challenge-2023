package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"synapsis-challange/helper"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	user                = os.Getenv("DB_USERNAME")
	password            = os.Getenv("DB_PASSWORD")
	hostname            = os.Getenv("DB_HOSTNAME")
	port                = os.Getenv("DB_PORT")
	databaseName        = os.Getenv("DB_DATABASE")
	mongoPoolMin, _     = strconv.Atoi(os.Getenv("DB_POOL_MIN"))
	mongoPoolMax, _     = strconv.Atoi(os.Getenv("DB_POOL_MAX"))
	mongoMaxIdleTime, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME_SECOND"))
	setTimeout, _       = strconv.Atoi(os.Getenv("DB_TIMEOUT"))
)

func NewDB() *mongo.Database {
	ctx, cancel := NewDBContext()
	defer cancel()

	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", user, password, hostname, port, databaseName)

	option := options.Client().
		ApplyURI(dsn).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	helper.PanicIfError(err)

	err = client.Connect(ctx)
	helper.PanicIfError(err)

	err = client.Ping(ctx, readpref.Primary())
	helper.PanicIfError(err)

	log.Println("Database connected")

	database := client.Database(databaseName)

	return database
}

func NewDBContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(setTimeout)*time.Second)
}

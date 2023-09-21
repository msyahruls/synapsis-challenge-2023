package helper

import (
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Pagging(page, limit string, sort int) *options.FindOptions {
	options := options.Find()
	if page != "" {
		if limit == "" {
			limit = os.Getenv("DEFAULT_LIMIT")
		}
		intPage, _ := strconv.ParseInt(page, 10, 32)
		intLimit, _ := strconv.ParseInt(limit, 10, 32)

		if intPage == 1 {
			options.SetSkip(0)
			options.SetLimit(intLimit)
		} else {
			options.SetSkip((intPage - 1) * intLimit)
			options.SetLimit(intLimit)
		}
	} else if limit != "" {
		intLimit, _ := strconv.ParseInt(limit, 10, 32)
		options.SetLimit(intLimit)
	}
	return options
}

func PaggingAggregat(page, limit string, pipeline *mongo.Pipeline) {
	if page != "" {
		if limit == "" {
			limit = os.Getenv("DEFAULT_LIMIT")
		}
		intPage, _ := strconv.ParseInt(page, 10, 32)
		intLimit, _ := strconv.ParseInt(limit, 10, 32)

		if intPage == 1 {
			skipStage := bson.D{{Key: "$skip", Value: 0}}
			limitStage := bson.D{{Key: "$limit", Value: intLimit}}
			*pipeline = append(*pipeline, skipStage, limitStage)
		} else {
			skipStage := bson.D{{Key: "$skip", Value: (intPage - 1) * intLimit}}
			limitStage := bson.D{{Key: "$limit", Value: intLimit}}
			*pipeline = append(*pipeline, skipStage, limitStage)
		}

	} else if limit != "" {
		intLimit, _ := strconv.ParseInt(limit, 10, 32)
		limitStage := bson.D{{Key: "$limit", Value: intLimit}}
		*pipeline = append(*pipeline, limitStage)
	}
}

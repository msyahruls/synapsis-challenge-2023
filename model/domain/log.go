package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Logs struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Actor     string             `json:"actor" bson:"actor"`
	Action    string             `json:"action" bson:"action"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}

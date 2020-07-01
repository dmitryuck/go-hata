package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserID  primitive.ObjectID `bson:"userId"`
	Picture string             `bson:"picture"`
	Message string             `bson:"message"`
	Created time.Time          `bson:"created"`
}

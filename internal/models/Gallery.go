package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gallery struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserID  primitive.ObjectID `bson:"userId"`
	Created time.Time          `bson:"created"`
}

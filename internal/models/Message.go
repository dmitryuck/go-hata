package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	SenderID primitive.ObjectID `bson:"senderId"`
	Text     string             `bson:"text"`
	IsRead   bool               `bson:"isRead"`
	Created  time.Time          `bson:"created"`
}

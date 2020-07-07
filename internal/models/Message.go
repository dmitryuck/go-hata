package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	SenderID primitive.ObjectID `json:"senderId"`
	Text     string             `json:"text"`
	IsRead   bool               `json:"isRead"`
	Created  time.Time          `json:"created"`
}

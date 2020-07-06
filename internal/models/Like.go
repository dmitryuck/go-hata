package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Like struct {
	IsOpen bool               `json:"isOpen"`
	UserID primitive.ObjectID `json:"userId"`
}

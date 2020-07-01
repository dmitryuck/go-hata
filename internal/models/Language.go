package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Language struct {
	ID   primitive.ObjectID `bson:"_id"`
	Code string             `bson:"code"`
	Name string             `bson:"name"`
}

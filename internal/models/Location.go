package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID      primitive.ObjectID `bson:"_id"`
	Code    string             `bson:"code"`
	Country string             `bson:"country"`
	Regions interface{}        `bson:"regions"`
}

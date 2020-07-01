package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Sex      string             `bson:"sex"`
	About    string             `bson:"about"`
	Birth    time.Time          `bson:"birth"`
	Photos   []string           `bson:"photos"`
	Likes    string             `bson:"likes"`
	Guests   string             `bson:"guests"`
	Location []string           `bson:"location"`
	Language string             `bson:"language"`
	Money    int                `bson:"money"`
	Active   bool               `bson:"active"`
	DeviceID string             `bson:"deviceId"`
	Updated  time.Time          `bson:"updated"`
	Created  time.Time          `bson:"created"`
}

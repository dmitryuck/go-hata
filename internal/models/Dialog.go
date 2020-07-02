package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dialog struct {
	ID        primitive.ObjectID   `bson:"_id"`
	MemberIDs []primitive.ObjectID `bson:"memberIds"`
	Users     []*User
	Messages  []Message            `bson:"messages"`
	DeleteIDs []primitive.ObjectID `bson:"deleteIds"`
	Updated   time.Time            `bson:"updated"`
	Created   time.Time            `bson:"created"`
}

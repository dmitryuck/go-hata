package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dialog struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	MemberIDs []primitive.ObjectID `bson:"memberIds"`
	Users     []*User              `json:"users"`
	Messages  []*Message           `bson:"messages"`
	DeleteIDs []primitive.ObjectID `bson:"deleteIds"`
	Updated   time.Time            `bson:"updated"`
	Created   time.Time            `bson:"created"`
}

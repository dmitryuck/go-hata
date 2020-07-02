package aggregates

import (
	"context"
	"project/internal/db"
	"project/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AggregateDialogs(profileIDStr string, userIDStr string) ([]*models.Dialog, error) {
	context := context.TODO()

	var userIDMatch bson.D = bson.D{{}}

	profileID, _ := primitive.ObjectIDFromHex(profileIDStr)

	if userIDStr != "" {
		userID, _ := primitive.ObjectIDFromHex(userIDStr)

		userIDMatch = bson.D{{"memberIds", bson.M{"$in": []primitive.ObjectID{userID}}}}
	}

	matchStage := bson.D{
		{"$match", bson.M{
			"$and": bson.A{
				bson.M{"memberIds": bson.M{"$in": []primitive.ObjectID{profileID}}},
				userIDMatch,
				bson.M{"deleteIds": bson.M{"$nin": []primitive.ObjectID{profileID}}},
			},
		}},
	}

	lookupState := bson.D{
		{"$lookup", bson.M{
			"from":         "users",
			"localField":   "memberIds",
			"foreignField": "_id",
			"as":           "users",
		}},
	}

	collection := db.Instance.Database.Collection("dialogs")

	curr, err := collection.Aggregate(context, mongo.Pipeline{matchStage, lookupState})
	if err != nil {
		return nil, err
	}

	var dialogs []*models.Dialog

	if err = curr.All(context, &dialogs); err != nil {
		return nil, err
	}

	return dialogs, nil
}

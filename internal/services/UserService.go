package services

import (
	"context"
	"encoding/json"
	"time"

	"project/internal/db"
	"project/internal/models"
	"project/internal/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchProfile feth profile
func FetchProfile(deviceIDStr string) (*response.UserResponse, error) {
	usersCollection := db.Instance.Database.Collection("users")

	var profile *models.User

	context := context.TODO()

	usersCollection.FindOne(context, bson.M{"deviceId": deviceIDStr}).Decode(&profile)

	likes, err := json.Marshal(&models.Likes{
		YesLikes:    []string{},
		NoLikes:     []string{},
		SmbLikes:    []models.Like{},
		SympatSet:   []string{},
		SympatFetch: "",
	})
	if err != nil {

	}

	guests, err := json.Marshal(&models.Guests{
		MyGuests: []models.Like{},
	})
	if err != nil {
		return nil, err
	}

	if profile == nil {
		profile = &models.User{
			Name:     "",
			Birth:    time.Now(),
			Photos:   []string{},
			Likes:    string(likes),
			Guests:   string(guests),
			Location: []string{},
			Money:    500,
			Active:   true,
			DeviceID: deviceIDStr,
			Updated:  time.Now(),
			Created:  time.Now(),
		}

		createdUser, err := usersCollection.InsertOne(context, profile)
		if err != nil {
			return nil, err
		}

		usersCollection.FindOne(context, bson.M{"_id": createdUser.InsertedID}).Decode(&profile)
	} else {
		after := options.After

		usersCollection.FindOneAndUpdate(context, bson.M{"_id": profile.ID}, bson.M{"updated": time.Now()}, &options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
		}).Decode(&profile)
	}

	return response.MakeUserResponse(profile), nil
}

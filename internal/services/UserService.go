package services

import (
	"context"
	"encoding/json"
	"time"

	"project/internal/db"
	"project/internal/models"
	"project/internal/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return nil, err
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

// UpdateProfile update user
func UpdateProfile(profileIDStr string, body *models.User) (*response.UserResponse, error) {
	profileID, _ := primitive.ObjectIDFromHex(profileIDStr)

	usersCollection := db.Instance.Database.Collection("users")

	var profile *models.User

	after := options.After

	usersCollection.FindOneAndUpdate(context.TODO(), bson.M{"_id": profileID}, body, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}).Decode(&profile)

	return response.MakeUserResponse(profile), nil
}

type ProfileCounts struct {
	Profile *response.UserResponse `json:"profile"`
	Counts  *models.PanelCounts    `json:"counts"`
}

// FetchProfileCounts fetch profile and counts
func FetchProfileCounts(profileIDStr string) (*ProfileCounts, error) {
	panelCounts := &models.PanelCounts{}

	profileCounts := &ProfileCounts{
		Profile: nil,
		Counts:  panelCounts,
	}

	return profileCounts, nil
}

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
	profileID, _ := primitive.ObjectIDFromHex(profileIDStr)

	panelCounts := &models.PanelCounts{}

	profileCounts := &ProfileCounts{
		Profile: nil,
		Counts:  panelCounts,
	}

	usersCollection := db.Instance.Database.Collection("users")

	var profile *models.User

	context := context.TODO()

	usersCollection.FindOne(context, bson.M{"_id": profileID}).Decode(&profile)

	if profile != nil {
		dialogs, err := FetchDialogs(profileIDStr)
		if err != nil {
			return nil, err
		}

		for _, dialog := range dialogs {
			panelCounts.Dialogs += dialog.UnreadCount
		}

		var profileLikeObj models.Likes

		if err := json.Unmarshal([]byte(profile.Likes), &profileLikeObj); err != nil {
			return nil, err
		}

		var likeUserIDs []primitive.ObjectID

		for _, like := range profileLikeObj.SmbLikes {
			if !like.IsOpen {
				likeUserIDs = append(likeUserIDs, like.UserID)
			}
		}

		var likeUsers []models.User

		curLikes, err := usersCollection.Find(context, bson.M{"_id": bson.M{"$in": likeUserIDs}})
		if err != nil {
			return nil, err
		}

		for curLikes.Next(context) {
			var user models.User

			if err := curLikes.Decode(&user); err != nil {
				return nil, err
			}

			likeUsers = append(likeUsers, user)
		}

		if err := curLikes.Err(); err != nil {
			return nil, err
		}

		curLikes.Close(context)

		panelCounts.Sympats = len(likeUsers)

		var profileGuestsObject models.Guests

		if err := json.Unmarshal([]byte(profile.Guests), &profileGuestsObject); err != nil {
			return nil, err
		}

		var guestUserIDs []primitive.ObjectID

		for _, like := range profileGuestsObject.MyGuests {
			if !like.IsOpen {
				guestUserIDs = append(guestUserIDs, like.UserID)
			}
		}

		var guestUsers []models.User

		curGuests, err := usersCollection.Find(context, bson.M{"_id": bson.M{"$in": guestUserIDs}})
		if err != nil {
			return nil, err
		}

		for curGuests.Next(context) {
			var guest models.User

			if err := curGuests.Decode(&guest); err != nil {
				return nil, err
			}

			guestUsers = append(guestUsers, guest)
		}

		if err := curGuests.Err(); err != nil {
			return nil, err
		}

		curGuests.Close(context)

		panelCounts.Guests = len(guestUsers)
	}

	return profileCounts, nil
}

// FetchPeopleUsers fetch people options
func FetchPeopleUsers(profileIDStr string, filterOptionsStr string, usersOffsetStr string) ([]*response.UserResponse, error) {
	profileID, _ := primitive.ObjectIDFromHex(profileIDStr)

	usersCollection := db.Instance.Database.Collection("users")

	var profile *models.User

	context := context.TODO()

	usersCollection.FindOne(context, bson.M{"_id": profileID}).Decode(&profile)

	locationsCollection := db.Instance.Database.Collection("locations")

	var locations []models.Location

	curLocations, err := locationsCollection.Find(context, bson.M{})
	if err != nil {
		return nil, err
	}

	for curLocations.Next(context) {
		var location models.Location

		if err := curLocations.Decode(&location); err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	if err := curLocations.Err(); err != nil {
		return nil, err
	}

	curLocations.Close(context)

	var targetSex string

	if profile.Sex == "male" {
		targetSex = "female"
	}

	if profile.Sex == "female" {
		targetSex = "male"
	}

	usersMatch := bson.M{
		"_id":    bson.M{"$ne": profileID},
		"photos": bson.M{"$exists": true, "$ne": bson.A{}},
	}

	if targetSex != "" {
		usersMatch["sex"] = targetSex
	}

	var users []models.User

	curUsers, err := usersCollection.Find(context, usersMatch)
	if err != nil {
		return nil, err
	}

	for curUsers.Next(context) {
		var user models.User

		if err := curUsers.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := curUsers.Err(); err != nil {
		return nil, err
	}

	curUsers.Close(context)

	var filterOptions models.FilterOptions

	if err := json.Unmarshal([]byte(filterOptionsStr), &filterOptions); err != nil {
		return nil, err
	}

	var usersResponse []*response.UserResponse

	for _, user := range users {
		if ApplyFilters(&user, filterOptions) {
			usersResponse = append(usersResponse, response.MakeUserResponse(&user))
		}
	}

	return usersResponse, nil
}

// ApplyFilters apply filters to user
func ApplyFilters(user *models.User, filterOptions models.FilterOptions) bool {
	if filterOptions.MinAge != 0 && filterOptions.MaxAge != 0 {

	}

	return true
}

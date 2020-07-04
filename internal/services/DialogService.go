package services

import (
	"context"
	"project/internal/aggregates"
	"project/internal/db"
	"project/internal/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FetchDialogs fetch dialogs
func FetchDialogs(profileIDStr string) ([]*response.DialogResponse, error) {
	/*context := context.TODO()

	collection := db.Instance.Database.Collection("dialogs")

	var dialogs []*models.Dialog

	cur, err := collection.Find(context, bson.M{})
	if err != nil {
		return nil, err
	}

	for cur.Next(context) {
		var d models.Dialog

		err := cur.Decode(&d)
		if err != nil {
			return nil, err
		}

		dialogs = append(dialogs, &d)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context)*/

	dialogs, err := aggregates.AggregateDialogs(profileIDStr, "")
	if err != nil {
		return nil, err
	}

	var dialogsResponse []*response.DialogResponse

	for _, dialog := range dialogs {
		dialogsResponse = append(dialogsResponse, response.MakeDialogResponse(dialog, profileIDStr))
	}

	return dialogsResponse, nil
}

// LoadDialog load dialog
func LoadDialog(profileIDStr string, userIDStr string) (*response.DialogResponse, error) {
	dialogs, err := aggregates.AggregateDialogs(profileIDStr, userIDStr)
	if err != nil {
		return nil, err
	}

	return response.MakeDialogResponse(dialogs[0], profileIDStr), nil
}

// CreateDialog create new dialog
func CreateDialog(profileIDStr string, userIDStr string) (*response.DialogResponse, error) {
	collection := db.Instance.Database.Collection("dialogs")

	profileID, _ := primitive.ObjectIDFromHex(profileIDStr)
	userID, _ := primitive.ObjectIDFromHex(userIDStr)

	_, err := collection.InsertOne(context.TODO(), bson.M{"memberIds": bson.A{profileID, userID}})
	if err != nil {
		return nil, err
	}

	dialogs, err := aggregates.AggregateDialogs(profileIDStr, userIDStr)
	if err != nil {
		return nil, err
	}

	return response.MakeDialogResponse(dialogs[0], profileIDStr), nil
}

package services

import (
	"context"
	"project/internal/aggregates"
	"project/internal/db"
	"project/internal/models"
	"project/internal/response"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchDialogs fetch dialogs
func FetchDialogs(profileIDStr string) ([]*response.DialogResponse, error) {
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
	profileID, _ := primitive.ObjectIDFromHex(profileIDStr)
	userID, _ := primitive.ObjectIDFromHex(userIDStr)

	dialogsCollection := db.Instance.Database.Collection("dialogs")

	_, err := dialogsCollection.InsertOne(context.TODO(), bson.M{"memberIds": bson.A{profileID, userID}})
	if err != nil {
		return nil, err
	}

	dialogs, err := aggregates.AggregateDialogs(profileIDStr, userIDStr)
	if err != nil {
		return nil, err
	}

	return response.MakeDialogResponse(dialogs[0], profileIDStr), nil
}

// SendMessage send message
func SendMessage(profileIDStr string, dialogIDStr string, text string) (*response.DialogResponse, error) {
	profileID, _ := primitive.ObjectIDFromHex(profileIDStr)
	dialogID, _ := primitive.ObjectIDFromHex(dialogIDStr)

	context := context.TODO()

	dialogsCollection := db.Instance.Database.Collection("dialogs")

	var dialog models.Dialog

	dialogsCollection.FindOne(context, bson.M{"_id": dialogID}).Decode(&dialog)

	after := options.After

	var updatedDialog models.Dialog

	dialogsCollection.FindOneAndUpdate(context, bson.M{"_id": dialogID}, bson.M{
		"messages": append(dialog.Messages, &models.Message{
			SenderID: profileID,
			Text:     text,
			IsRead:   false,
			Created:  time.Now(),
		}),
		"deleteIds": bson.A{},
	}, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}).Decode(&updatedDialog)

	dialogs, err := aggregates.AggregateDialogs(updatedDialog.MemberIDs[0].String(), updatedDialog.MemberIDs[1].String())
	if err != nil {
		return nil, err
	}

	return response.MakeDialogResponse(dialogs[0], profileIDStr), nil
}

// DeleteDialog delete dialog
func DeleteDialog(profileIDStr string, dialogIDStr string) (bool, error) {
	profileID, _ := primitive.ObjectIDFromHex(profileIDStr)
	dialogID, _ := primitive.ObjectIDFromHex(dialogIDStr)

	context := context.TODO()

	dialogsCollection := db.Instance.Database.Collection("dialogs")

	var dialog models.Dialog

	dialogsCollection.FindOne(context, bson.M{"_id": dialogID}).Decode(&dialog)

	if dialog.DeleteIDs != nil && len(dialog.DeleteIDs) > 0 {
		dialogsCollection.FindOneAndDelete(context, bson.M{"_id": dialogID})
	} else {
		dialogsCollection.FindOneAndUpdate(context, bson.M{"_id": dialogID}, bson.M{
			"deleteIds": bson.A{profileID},
		})
	}

	return true, nil
}

// UpdateDialog update dialog
func UpdateDialog(dialogIDStr string, body *models.Dialog) (*response.DialogResponse, error) {
	dialogID, _ := primitive.ObjectIDFromHex(dialogIDStr)

	dialogsCollection := db.Instance.Database.Collection("dialogs")

	after := options.After

	var updatedDialog models.Dialog

	dialogsCollection.FindOneAndUpdate(context.TODO(), bson.M{"_id": dialogID}, body, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}).Decode(&updatedDialog)

	dialogs, err := aggregates.AggregateDialogs(updatedDialog.MemberIDs[0].String(), updatedDialog.MemberIDs[1].String())
	if err != nil {
		return nil, err
	}

	return response.MakeDialogResponse(dialogs[0], ""), nil
}

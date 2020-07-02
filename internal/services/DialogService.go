package services

import (
	"project/internal/aggregates"
	"project/internal/response"
)

// FetchDialogs fetch dialogs
func FetchDialogs(profileID string) ([]*response.DialogResponse, error) {
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

	dialogs, err := aggregates.AggregateDialogs(profileID, "")
	if err != nil {
		return nil, err
	}

	var dialogsResponse []*response.DialogResponse

	for _, dialog := range dialogs {
		dialogsResponse = append(dialogsResponse, response.DialogResponse.Make(response.DialogResponse{}, dialog))
	}

	return dialogsResponse, nil
}

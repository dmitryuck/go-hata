package services

import (
	"context"

	"project/internal/db"
	"project/internal/logger"
	"project/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

type DialogService struct {
}

// FetchDialogs fetch dialogs
func (s DialogService) FetchDialogs() models.Dialog {
	logger.Instance.LogInfo("YESS")

	var dialog models.Dialog

	collection := db.Instance.Database.Collection("dialogs")

	collection.FindOne(context.TODO(), bson.M{}).Decode(&dialog)

	logger.Instance.LogInfo("FFFF")

	return dialog
}

package response

import (
	"time"

	"project/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DialogResponse struct {
	ID          primitive.ObjectID   `json:"_id"`
	MemberIDs   []primitive.ObjectID `json:"memberIds"`
	Users       []*UserResponse      `json:"users"`
	DeleteIDs   []primitive.ObjectID `json:"deleteIds"`
	Messages    []*models.Message    `json:"messages"`
	UnreadCount int                  `json:"unreadCount"`
	Updated     time.Time            `json:"updated"`
	Created     time.Time            `json:"created"`
}

// MakeDialogResponse DialogResponse
func MakeDialogResponse(dialog *models.Dialog, profileIDStr string) *DialogResponse {
	profileID, _ := primitive.ObjectIDFromHex(profileIDStr)

	var users []*UserResponse

	for _, user := range dialog.Users {
		users = append(users, MakeUserResponse(user))
	}

	var unreadMessages []*models.Message

	for _, message := range dialog.Messages {
		if message.SenderID != profileID && !message.IsRead {
			unreadMessages = append(unreadMessages, message)
		}
	}

	return &DialogResponse{
		ID:          dialog.ID,
		MemberIDs:   dialog.MemberIDs,
		Users:       users,
		DeleteIDs:   dialog.DeleteIDs,
		Messages:    dialog.Messages,
		UnreadCount: len(unreadMessages),
		Updated:     dialog.Updated,
		Created:     dialog.Created,
	}
}

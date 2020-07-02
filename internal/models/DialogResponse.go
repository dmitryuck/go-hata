package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DialogResponse struct {
	ID          primitive.ObjectID   `json:"id"`
	MemberIDs   []primitive.ObjectID `json:"memberIds"`
	Users       []UserResponse       `json:"users"`
	DeleteIDs   []primitive.ObjectID `json:"deleteIds"`
	Messages    []Message            `json:"messages"`
	UnreadCount int                  `json:"unreadCount"`
	Updated     time.Time            `json:"updated"`
	Created     time.Time            `json:"created"`
}

// Make DialogResponse
func (d DialogResponse) Make(dialog *Dialog) *DialogResponse {
	//const users = dialog.users.map((user: User) => UserResponse.make(user));
	//const unreadCount = dialog.messages.filter(message => !isIdsEqual(message.senderId, profileIdObj) && !message.isRead).length;

	return &DialogResponse{
		ID:        dialog.ID,
		MemberIDs: dialog.MemberIDs,
		//Users: users,
		DeleteIDs: dialog.DeleteIDs,
		Messages:  dialog.Messages,
		//UnreadCount: unreadCount,
		Updated: dialog.Updated,
		Created: dialog.Created,
	}
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Sex      string             `json:"sex"`
	About    string             `json:"about"`
	Birth    time.Time          `json:"birth"`
	Photos   []string           `json:"photos"`
	Likes    string             `json:"likes"`
	Guests   string             `json:"guests"`
	City     string             `json:"city"`
	Country  string             `json:"country"`
	Location []string           `json:"location"`
	Language string             `json:"language"`
	Money    int                `json:"money"`
	Active   bool               `json:"active"`
	Updated  time.Time          `json:"updated"`
	Created  time.Time          `json:"created"`
}

func (u UserResponse) Make(user *User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Sex:      user.Sex,
		About:    user.About,
		Birth:    user.Birth,
		Photos:   user.Photos,
		Likes:    user.Likes,
		Guests:   user.Guests,
		City:     user.City,
		Country:  user.Country,
		Location: user.Location,
		Language: user.Language,
		Money:    user.Money,
		Active:   user.Active,
		Updated:  user.Updated,
		Created:  user.Created,
	}
}

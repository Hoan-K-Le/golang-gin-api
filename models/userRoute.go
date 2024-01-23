package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


type User struct {
	ID primitive.ObjectID `bson:"_id"`
	Username *string `json:"username" validate:"required", min=2,max=100`
	Email *string `json:"email" validate:"required", min=2,max=100`
	Password *string `json:"password" validate:"required", min=2,max=100`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at time.Time	`json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Cart []CartItem `json:"cart"`
}
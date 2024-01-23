package user

import (
	"github.com/Hoan-K-Le/golang-gin-api-ecom/structure/cart"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


type User struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	Username string `json:"username, omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required"`
	Password string `json:"password, omitempty" validate:"required"`
	Created_at time.time
	Updated_at time.Time
	Cart []cart.CartItem `json:"cart, omitempty" validate:"required"`
}
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string `json:"name"`
	Description string `json:"description"`
	Quantity int64 `json:"quantity"`
	Category string `json:"category"`
	ImageUrl string `json:"imageurl"`
}

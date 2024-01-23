package user

import "github.com/Hoan-K-Le/golang-gin-api-ecom/structure/cart"

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Cart []cart.CartItem `json:"cart"`
}
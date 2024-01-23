package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-lang-api-ecom/structure/cart"
	"github.com/go-lang-api-ecom/structure/user"
)



func main() {
	cartItem := cart.CartItem{
		ProductID: 123,
		Quantity: 1,
		Category: "Shirt"
	}

	user := user.User{
		ID: 1,
		Username: "test",
		Password: "test",
		Cart: []cart.CartItem{cartItem},
	}
	fmt.Printf("User: %+v\n",)
}
package main

import (
	// "net/http"

	"github.com/Hoan-K-Le/golang-gin-api-ecom/configs"
	"github.com/gin-gonic/gin"

	// "github.com/Hoan-K-Le/golang-gin-api-ecom/structure/cart"
	// "github.com/Hoan-K-Le/golang-gin-api-ecom/structure/user"
	"github.com/Hoan-K-Le/golang-gin-api-ecom/routes"

)



func main() {
	router := gin.Default();
	configs.ConnectDB()
	routes.UserRoute(router)
	router.Run("localhost:8000")
}
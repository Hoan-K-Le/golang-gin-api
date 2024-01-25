package main

import (
	  "github.com/gin-contrib/cors"
	   "log"
    "os"
    "github.com/joho/godotenv"
	"github.com/Hoan-K-Le/golang-gin-api-ecom/configs"
	"github.com/gin-gonic/gin"
	middleware "github.com/Hoan-K-Le/golang-gin-api-ecom/middleware"
	"github.com/Hoan-K-Le/golang-gin-api-ecom/routes"
	"github.com/Hoan-K-Le/golang-gin-api-ecom/seed"

)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default();
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET","PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
}))
 client := configs.ConnectDB()
	configs.ConnectDB()
	seed.SeedProducts(client)
	routes.UserRoute(router)
	router.Use(middleware.Authentication())
	router.Run(":" + os.Getenv("PORT"))
}
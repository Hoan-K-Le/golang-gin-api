package main

import (
	   "log"
    "os"
    "github.com/joho/godotenv"
	"github.com/Hoan-K-Le/golang-gin-api-ecom/configs"
	"github.com/gin-gonic/gin"
	middleware "github.com/Hoan-K-Le/golang-gin-api-ecom/middleware"
	"github.com/Hoan-K-Le/golang-gin-api-ecom/routes"

)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default();
	configs.ConnectDB()
	routes.UserRoute(router)
	router.Use(middleware.Authentication())

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})
    router.GET("/api-2", func(c *gin.Context) {
        c.JSON(200, gin.H{"success": "Access granted for api-2"})
    })
	router.Run(":" + os.Getenv("PORT"))
}
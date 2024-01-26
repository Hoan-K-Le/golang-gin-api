package controllers

import (
	"fmt"
	   "time"
		 "github.com/Hoan-K-Le/golang-gin-api-ecom/models"
		 "context"
		 "github.com/gin-gonic/gin"
		 "net/http"
		 "go.mongodb.org/mongo-driver/bson"
		 "github.com/Hoan-K-Le/golang-gin-api-ecom/configs"
)
// create products
func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var product models.Product
		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// insert the product into the database
		collection := configs.Client.Database("golangapi").Collection("products")

		result, err := collection.InsertOne(ctx, product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		fmt.Println(product)
		c.JSON(http.StatusOK, gin.H{"result":result})
	}
}

// GET /product
func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		// initialize a slice
		 var products []models.Product

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// access the products collection
		collection := configs.Client.Database("golangapi").Collection("products")

		// query the collection
		cursor, err := collection.Find(ctx,bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var product models.Product

			err := cursor.Decode(&product)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			products = append(products, product)
		}

		  if err := cursor.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

				c.JSON(http.StatusOK, products)
	}
}
// GET /product/:id

// edit products

// delete products

// filter products

// add product to user's cart
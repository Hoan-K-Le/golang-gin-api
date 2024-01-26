package main

import (
	"fmt"
	"github.com/Hoan-K-Le/golang-gin-api-ecom/models"
  "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"context"
)

var products = []models.Product {
	{ Name:"Classic T-Shirt", Description:"Classic white t-shirt", Quantity: 4, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"White T-Shirt", Description:"Classic white t-shirt", Quantity: 8, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"Black T-Shirt", Description:"Classic white t-shirt", Quantity: 7, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"Gray T-Shirt", Description:"Classic white t-shirt", Quantity: 2, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"Pink T-Shirt", Description:"Classic white t-shirt", Quantity: 10, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"Purple T-Shirt", Description:"Classic white t-shirt", Quantity: 12, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"Mix T-Shirt", Description:"Classic white t-shirt", Quantity: 2, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"Yellow T-Shirt", Description:"Classic white t-shirt", Quantity: 1, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"Brown T-Shirt", Description:"Classic white t-shirt", Quantity: 8, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"Red T-Shirt", Description:"Classic white t-shirt", Quantity: 4, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"},
	{ Name:"Classic Pants", Description:"Classic tye dyed pants", Quantity: 7, Category: "pants", ImageUrl: "https://placeholder.com/200/400"},
	{ Name:"Red Pants", Description:"Classic tye dyed pants", Quantity: 5, Category: "pants", ImageUrl: "https://placeholder.com/200/400"},
	{ Name:"Blue Pants", Description:"Classic tye dyed pants", Quantity: 3, Category: "pants", ImageUrl: "https://placeholder.com/200/400"},
	{ Name:"Black Pants", Description:"Classic tye dyed pants", Quantity: 7, Category: "pants", ImageUrl: "https://placeholder.com/200/400"},
	{ Name:"Green Pants", Description:"Classic tye dyed pants", Quantity: 8, Category: "pants", ImageUrl: "https://placeholder.com/200/400"},
	{ Name:"Yellow Pants", Description:"Classic tye dyed pants", Quantity: 10, Category: "pants", ImageUrl: "https://placeholder.com/200/400"},
	{ Name:"Aqua Pants", Description:"Classic tye dyed pants", Quantity: 1, Category: "pants", ImageUrl: "https://placeholder.com/200/400"},
}


func SeedProducts(client *mongo.Client) {
	// grab the collection from the database
    collection := client.Database("golangapi").Collection("products")

    // Check if the collection is empty, if not then print out that it has already been created
    count, err := collection.CountDocuments(context.Background(), bson.D{{}})
    if err != nil {
        log.Fatal(err)
    }
    if count > 0 {
        fmt.Println("Products already seeded")
        return // Collection is already seeded
    }
		// created a documents slice to insert to database later
	var documents []interface{}
	for _, product := range products {
		documents = append(documents, product)
	}

	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		log.Fatal("Failed to insert products",err)
	}
}
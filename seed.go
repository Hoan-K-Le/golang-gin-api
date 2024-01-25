package seed
import (
	"github.com/Hoan-K-Le/golang-gin-api-ecom/models"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"context"
)

var products = []models.Product {
	{ID: "1", Name:"Classic T-Shirt", Description:"Classic white t-shirt", Quantity: 4, Category: "shirt", ImageUrl: "https://placeholder.com/200/300"}
	{ID: "2", Name:"Classic Pants", Description:"Classic tye dyed pants", Quantity: 7, Category: "pants", ImageUrl: "https://placeholder.com/200/400"}
}


func SeedProducts(client *mongo.Client) {
collection := client.Database("golangapi").Collection("product")

    // Check if the collection is empty
    count, err := collection.CountDocuments(context.Background(), bson.D{{}})
    if err != nil {
        log.Fatal(err)
    }
    if count > 0 {
        fmt.Println("Products already seeded")
        return // Collection is already seeded
    }
	var documents []interface{}
	for _, product := range products {
		documents = append(documents, product)
	}

	_, err := collection.InsertMany(context.Background(), documents)
	if err != nil {
		log.Fatal("Failed to insert products",err)
	}
}
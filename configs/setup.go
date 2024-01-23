package configs

import (
    "context"
    "fmt"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
// function that connects to the mongo database and returns a mongo.Client pointer
func ConnectDB() *mongo.Client  {
    // mongo.NewClient creates a new MongoDB client and options... sets the MongoDB URI from .env using the helper EnvMongoURI
    client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
    if err != nil {
        log.Fatal(err)
    }
    // Creates a context with a 10-second timeout, ensuring operations don't hang indefinitely 
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    // defer cancel() makes sure the context is correctly canceled when function exits
    defer cancel()
    // connecting to MongoDB
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    // Pings the mongo database to verify if its active otherwise it logs the error and terminates
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
}

//Client instance(declared and initialized)
var Client *mongo.Client = ConnectDB()

//getting database collections
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("golangapi").Collection(collectionName)
    return collection
}

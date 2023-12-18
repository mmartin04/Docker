package db

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "time"
)

var Client *mongo.Client

func Connect() {
    uri := "mongodb://root:example@mongodb:27017" // Passen Sie die URI entsprechend an
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var err error
    Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    // Überprüfen Sie die Verbindung
    err = Client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
    return Client.Database("UserServiceDB").Collection(collectionName)
}

package router

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var db *mongo.Database

func ConnectToDB(mongoURI string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

    if client.Ping(context.TODO(), nil); err != nil {
        log.Fatal(err)
    }

	db = client.Database("tweets")
	log.Printf("Connected to database in: %s", mongoURI)
}

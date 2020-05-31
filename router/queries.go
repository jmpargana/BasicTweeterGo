package router

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetTweets(userIP string) ([]Tweet, error) {
	tweetsCollection := db.Collection("tweets")

	var tweets []Tweet
	cursor, err := tweetsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &tweets); err != nil {
		return nil, err
	}
	return tweets, nil
}

func performPost(obj interface{}, collectionName string) {
	tweetsCollection := db.Collection(collectionName)

	_, err := tweetsCollection.InsertOne(context.TODO(), obj)
	if err != nil {
		log.Printf("failed upload %v: %v", obj, err)
	}
}

package router

import (
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tweet struct {
	ID     primitive.ObjectID `bson:id,omitempty`
	Body   string             `bson:body,omitempty`
	UserIP string             `bson:user,omitempty`
	Date   time.Time          `bson:date,omitempty`
}

type User struct {
	ID        primitive.ObjectID   `bson:id,omitempty`
	UserIP    string               `bson:user,omitempty`
	Following []primitive.ObjectID `bson:following,omitempty`
}

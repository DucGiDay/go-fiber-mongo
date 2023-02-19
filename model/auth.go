package model

import (
	"context"
	"time"

	"github.com/hiiamtrong/go-fiber-restapi/config"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckInvalidCredentials(Username string, Password string) (User, error) {
	var MI config.MongoInstance = config.MI

	var user User
	collection := MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findResult := collection.FindOne(ctx, bson.M{"username": Username, "password": Password})
	if err := findResult.Err(); err != nil {
		return user, err
	}

	err := findResult.Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}

func CheckUserAlreadyExists(Username string) (User, error) {
	var MI config.MongoInstance = config.MI

	var user User
	collection := MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findResult := collection.FindOne(ctx, bson.M{"username": Username})
	if err := findResult.Err(); err != nil {
		return user, err
	}

	err := findResult.Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}

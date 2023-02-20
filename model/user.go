package model

import (
	"context"
	"log"
	"time"

	"github.com/hiiamtrong/go-fiber-restapi/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Age      int                `json:"age" bson:"age"`
}

func GetAllUsers() ([]User, error) {
	var MI config.MongoInstance = config.MI
	var users []User

	collection := MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}

	// for _, user := range users {
	// 	log.Println(user)
	// }

	return users, nil
}

func GetUser(userId primitive.ObjectID) (User, error) {
	var MI config.MongoInstance = config.MI

	var user User
	collection := MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findResult := collection.FindOne(ctx, bson.M{"_id": userId})
	if err := findResult.Err(); err != nil {
		return user, err
	}

	err := findResult.Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user User) (User, error) {
	var MI config.MongoInstance = config.MI
	collection := MI.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func UpdateUser(userId primitive.ObjectID, user User) (User, error) {
	var MI config.MongoInstance = config.MI

	userMap := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"age":      user.Age,
	}

	collection := MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$set": userMap})
	if err != nil {
		return user, err
	}
	return user, nil
}

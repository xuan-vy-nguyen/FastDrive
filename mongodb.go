package main

import (
	"context"
    "fmt"
    "log"

    _ "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/xuan-vy-nguyen/SE_Project01/database"
)

func addLoginDB(mail_ string, token_ string)(bool) {
    // Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println("Connected to MongoDB!")

	// insert to MongoDB
	collection := client.Database("app").Collection("LoginDB")

	newElement := database.LoginDB{
		Mail: mail_,
		Token: token_,
	}
	insertResult, err := collection.InsertOne(context.TODO(), newElement)
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println("Connection to MongoDB closed.")
	return false
}

func addSignUpDB(infor database.SignUpAccount) string {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return "server has something wrong"
	}
	fmt.Println("Connected to MongoDB!")

	// insert to MongoDB
	collection := client.Database("app").Collection("SignDB")

	newElement := infor
	insertResult, err := collection.InsertOne(context.TODO(), newElement)
	if err != nil {
		log.Fatal(err)
		return "server has something wrong"
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return "server has something wrong"
	}
	fmt.Println("Connection to MongoDB closed.")
	return ""
}

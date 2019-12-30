package database

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // You will be using this Trainer type later in the program
// type Trainer struct {
// 	Name string
// 	Age  int
// 	City string
// }

// func main_next() {
// 	// Set client options
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")

// 	// Insert something to DB
// 	collection := client.Database("test").Collection("trainers")
// 	ash := Trainer{"Ash", 10, "Pallet Town"}
// 	misty := Trainer{"Misty", 10, "Cerulean City"}
// 	brock := Trainer{"Brock", 15, "Pewter City"}

// 	insertResult, err := collection.InsertOne(context.TODO(), ash)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

// 	// insert many to DB
// 	trainers := []interface{}{misty, brock}

// 	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

// 	// Disconnect
// 	err = client.Disconnect(context.TODO())

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connection to MongoDB closed.")
// }

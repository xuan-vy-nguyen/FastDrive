package controllers

import (
	"context"
	"fmt"
	"log"

	datastruct "github.com/xuan-vy-nguyen/SE_Project01/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddOneImageDB insert one image to grid db
func AddOneImageDB(data []byte, mail string, filename string) bool { // return err
	clientOptions := options.Client().ApplyURI(MongoURI)
	// Connect to MongoDB
	clientOptions.SetMaxPoolSize(5)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println("Connected to ImageDB!")

	// insert Image to mongoDB
	collection := client.Database(Collection).Collection(ImageDB)

	newElement := datastruct.ImageDB{
		Name:  filename,
		Mail:  mail,
		Image: data,
	}
	insertResult, err := collection.InsertOne(context.TODO(), newElement)
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println("Inserted a single document: ", insertResult)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return true
	}

	fmt.Println("Connection to ImageDB closed.")

	return false
}

// GetOneImageDB is ok
func GetOneImageDB(fileName string, mail string) (datastruct.ImageDB, bool) { // return image and err
	var result datastruct.ImageDB
	clientOptions := options.Client().ApplyURI(MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return result, true
	}

	fmt.Println("Connected to ImageDB!")

	// insert Image to mongoDB
	collection := client.Database(Collection).Collection(ImageDB)

	// set filter
	filter := bson.D{
		primitive.E{Key: "name", Value: fileName},
		primitive.E{Key: "mail", Value: mail},
	}

	// this errCondition is use below
	errCondition := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("finding error ", errCondition)
	if errCondition != nil {
		return result, true
	}

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return result, true
	}
	fmt.Println("Connection to ImageDB closed.")

	// return
	return result, false
}

// GetAllNameUserImage return list of all user's images. Each user has a private list.
func GetAllNameUserImage(mail string) ([]string, bool) { // return list of name images and err
	var result []string
	clientOptions := options.Client().ApplyURI(MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return result, true
	}

	fmt.Println("Connected to ImageDB!")

	// insert Image to mongoDB
	collection := client.Database(Collection).Collection(ImageDB)

	// set filter
	filter := bson.D{primitive.E{Key: "mail", Value: mail}}

	// get list of name image
	cur, errCondition := collection.Find(context.Background(), filter)
	if errCondition != nil {
		fmt.Println("finding error ", errCondition)
		return result, true
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var ele datastruct.ImageDB
		err := cur.Decode(&ele)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...
		result = append(result, ele.Name)
	}
	if err := cur.Err(); err != nil {
		return result, true
	}

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return result, true
	}
	fmt.Println("Connection to ImageDB closed.")

	// return
	return result, false
}

// DeleteOneImageDB delete one image
func DeleteOneImageDB(fileName string, mail string) bool { // return err
	clientOptions := options.Client().ApplyURI(MongoURI)
	// Connect to MongoDB
	clientOptions.SetMaxPoolSize(5)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println("Connected to ImageDB!")

	// delete Image int mongoDB
	collection := client.Database(Collection).Collection(ImageDB)

	// set filter
	filter := bson.D{
		primitive.E{Key: "name", Value: fileName},
		primitive.E{Key: "mail", Value: mail},
	}

	// this errCondition is use below
	_, errDelete := collection.DeleteOne(context.TODO(), filter)
	if errDelete != nil {
		fmt.Println("finding error ", errDelete)
		return true
	}

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return true
	}

	fmt.Println("Connection to ImageDB closed.")

	return false
}

package dbactions

import (
	"context"
	"fmt"
	"log"

	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddOneImageDB insert one image to grid db
func AddOneImageDB(data []byte, email string, filename string) bool { // return err
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
		Email: email,
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
func GetOneImageDB(fileName string, jwtStr string) (datastruct.ImageDB, bool) { // return image and err
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

	// find userMail
	UserInfor, errGetMail := GetOneLoginDB(jwtStr)
	if errGetMail {
		return result, true
	}

	// set filter
	filter := bson.D{
		primitive.E{Key: "name", Value: fileName},
		primitive.E{Key: "email", Value: UserInfor.Mail},
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

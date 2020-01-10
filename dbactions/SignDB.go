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

// CheckAccInSignUpDB is ok
func CheckAccInSignUpDB(p datastruct.LoginAccount) int {
	clientOptions := options.Client().ApplyURI(MongoURI)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	fmt.Println("Connected to SignDB!")

	// find element in MongoDB
	collection := client.Database(Collection).Collection(SignDB)

	filter := bson.D{primitive.E{Key: "mail", Value: p.Mail}}

	var result datastruct.SignUpAccount
	// this errCondition is use below
	errCondition := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("finding error ", errCondition)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return 1
	}

	fmt.Println("Connection to SignDB closed.")

	// check errCondition (user wasnot declared)
	if errCondition != nil {
		return 4
	}
	// check password ok
	if result.Pass != p.Pass {
		return 0
	}
	return 2
}

// GetOneSignUpDB is ok
func GetOneSignUpDB(p string) (datastruct.SignUpAccount, bool) { // return err
	clientOptions := options.Client().ApplyURI(MongoURI)
	var result datastruct.SignUpAccount
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return result, true
	}
	fmt.Println("Connected to SignDB!")

	// find element in MongoDB
	collection := client.Database(Collection).Collection(SignDB)

	filter := bson.D{primitive.E{Key: "mail", Value: p}}

	// this errCondition is use below
	errCondition := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("finding error ", errCondition)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return result, true
	}

	fmt.Println("Connection to SignDB closed.")

	// check condition
	if errCondition != nil {
		return result, true
	}
	return result, false
}

// AddOneSignUpDB is ok
func AddOneSignUpDB(infor datastructSignUpAccount) string {
	clientOptions := options.Client().ApplyURI(MongoURI)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return "server has something wrong"
	}
	fmt.Println("Connected to SignDB!")

	// insert to MongoDB
	collection := client.Database(Collection).Collection(SignDB)

	newElement := infor
	insertResult, err := collection.InsertOne(context.TODO(), newElement)
	if err != nil {
		log.Fatal(err)
		return "server has something wrong"
	}
	fmt.Println("Inserted a single document: ", insertResult)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return "server has something wrong"
	}

	fmt.Println("Connection to SignDB closed.")
	return ""
}

// UpdateOneSignUpDB is ok
func UpdateOneSignUpDB(oldMail string, infor datastructSignUpAccount) bool { // err
	clientOptions := options.Client().ApplyURI(MongoURI)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println("Connected to SignDB!")

	// insert to MongoDB
	collection := client.Database(Collection).Collection(SignDB)

	newElement := infor
	filter := bson.D{primitive.E{Key: "mail", Value: oldMail}}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return true
	}
	res2, err := collection.InsertOne(context.TODO(), newElement)
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println("Update a single document: ", res2)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return true
	}

	fmt.Println("Connection to SignDB closed.")
	return false
}

// ExistInSignUpDB is ok
func ExistInSignUpDB(p string) bool { // return result
	clientOptions := options.Client().ApplyURI(MongoURI)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return true
	}
	fmt.Println("Connected to SignDB!")

	// find element in MongoDB
	collection := client.Database(Collection).Collection(SignDB)

	filter := bson.D{primitive.E{Key: "mail", Value: p}}

	var result datastruct.SignUpAccount

	// this errCondition is use below
	errCondition := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("finding error ", errCondition)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return true
	}

	fmt.Println("Connection to SignDB closed.")

	// check condition
	if errCondition != nil {
		return false
	}
	return true
}

package main

import (
	"context"
    "fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/xuan-vy-nguyen/SE_Project01/database"
)

var MongoURI = "mongodb+srv://xuanvy99:az1731999@cluster0-mzeio.mongodb.net/test?retryWrites=true&w=majority" // "mongodb://localhost:27017"
var Collection = "app"
var LoginDB = "LoginDB"
var SignDB = "SignDB"

func addLoginDB(mail_ string, token_ string)(bool) {	// return err
	clientOptions := options.Client().ApplyURI(MongoURI)	
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)	
	if err != nil {		
		log.Fatal(err)			 	
		return true		
	}
	fmt.Println("Connected to LoginDB!")

	// insert to MongoDB
	collection := client.Database(Collection).Collection(LoginDB)

	newElement := database.LoginDB{
		Mail: mail_,
		Token: token_,
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
	fmt.Println("Connection to LoginDB closed.")
	return false
}

func addSignUpDB(infor database.SignUpAccount) string {
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

func checkAccInSignUpDB(p database.LoginAccount)(int){
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

	var result database.SignUpAccount
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

func checkAccInLoginDB(p database.LoginAccount)(bool){	// return result 
	clientOptions := options.Client().ApplyURI(MongoURI)	
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)	
	if err != nil {		
		log.Fatal(err)			
		return true		
	}
	fmt.Println("Connected to LoginDB!")

	// find element in MongoDB
	collection := client.Database(Collection).Collection(LoginDB)

	filter := bson.D{primitive.E{Key: "mail", Value: p.Mail}}

	var result database.LoginDB

	// this errCondition is use below
	errCondition := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("finding error ", errCondition)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return true
	}
	
	fmt.Println("Connection to LoginDB closed.")
	
	// check condition
	if errCondition != nil {
		return false
	}
	return true
}

func existInSignUpDB(p string)(bool){	// return result
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

	var result database.SignUpAccount

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

func isInLoginDB(jwtStr string)(bool){	// return result
	clientOptions := options.Client().ApplyURI(MongoURI)	
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)	
	if err != nil {		
		log.Fatal(err)			
		return true		
	}
	fmt.Println("Connected to LoginDB!")

	// find element in MongoDB
	collection := client.Database(Collection).Collection(LoginDB)

	filter := bson.D{primitive.E{Key: "token", Value: jwtStr}}

	var result database.LoginDB

	// this errCondition is use below
	errCondition := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("finding error ", errCondition)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return false
	}
	
	fmt.Println("Connection to LoginDB closed.")
	
	// check condition
	if errCondition != nil {
		return false
	}
	return true
}

func removeInLoginDB(jwtStr string)(bool){	// return err
	clientOptions := options.Client().ApplyURI(MongoURI)	
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)	
	if err != nil {		
		log.Fatal(err)			
		return true		
	}

	// delete element in MongoDB
	fmt.Println("Connected to LoginDB!")

	collection := client.Database(Collection).Collection(LoginDB)

	filter := bson.D{primitive.E{Key: "token", Value: jwtStr}}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return true
	}
	
	fmt.Printf("Deleted %+v\n", deleteResult) 
	
	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return true
	}
	
	fmt.Println("Connection to LoginDB closed.")
	return false
}
package DBActions

import (
	"context"
    "fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/xuan-vy-nguyen/SE_Project01/DataStruct"
)

func GetOneSignUpDB(p string)(datastruct.SignUpAccount, bool){ // return err
	clientOptions := options.Client().ApplyURI(MongoURI)
	var result DataStructSignUpAccount	
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


func AddOneSignUpDB(infor DataStructSignUpAccount) string {
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

func UpdateOneSignUpDB(oldMail string, infor DataStructSignUpAccount) bool { // err
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


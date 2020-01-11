package dbactions

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddOneImageDB insert one image to grid db
func AddOneImageDB(data []byte, filename string) bool { // return err
	clientOptions := options.Client().ApplyURI(MongoURI)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return true
	}

	fmt.Println("Connected to ImageDB!")

	// insert Image to mongoDB
	bucket, err := gridfs.NewBucket(
		client.Database(Collection),
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename,
	)
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		log.Fatal(err)
		return true
	}
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return true
	}

	fmt.Println("Connection to ImageDB closed.")

	return false
}

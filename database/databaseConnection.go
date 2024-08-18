package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	envFile, _ := godotenv.Read(".env")

	mongoUser := envFile["MONGO_INITDB_ROOT_USERNAME"]
	mongoPass := envFile["MONGO_INITDB_ROOT_PASSWORD"]
	mongoHost := envFile["MONGO_HOST"]
	mongoPort := envFile["MONGO_PORT"]

	fmt.Printf("Username: %s, Password: %s, Host: %s, Port: %s\n", mongoUser, mongoPass, mongoHost, mongoPort)

	if mongoUser == "" || mongoPass == "" || mongoHost == "" || mongoPort == "" {
		log.Fatal("One or more required environment variables are missing")
	}

	MongoDb := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUser, mongoPass, mongoHost, mongoPort)
	fmt.Print(MongoDb)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}

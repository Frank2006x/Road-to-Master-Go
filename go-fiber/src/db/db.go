package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	_=godotenv.Load()

	mongoURI:= os.Getenv("MONGO_URL")
	dbName:= os.Getenv("DB_NAME")

	if (mongoURI == "" || dbName == "") {
		log.Fatal("Missing mongo env")
	}
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()


	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect mongoDB",err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping mongoDB",err)
	}
	log.Println("Connected to MongoDB!")
	DB = client
}

func GetCollection(collectionName string) *mongo.Collection {
	dbName:= os.Getenv("DB_NAME")
	return DB.Database(dbName).Collection(collectionName)
}
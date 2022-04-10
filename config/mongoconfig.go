package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BuildDBConfig() MongoCofig {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))

	clientOpts.Auth = &options.Credential{Username: username, Password: password}
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mongo")
	return MongoCofig{DB: client, DatabaseName: "bank"}
}

type MongoCofig struct {
	DB           *mongo.Client
	DatabaseName string
}

package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BuildDBConfig() MongoCofig {
	host := "localhost"
	port := 27017
	username := "root"
	password := "example"

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))

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

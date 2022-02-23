package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BuildDBConfig() MongoCofig {
	host := "localhost"   //TODO GET FROM ENVIROMENTS VARIABLES
	port := 27017         //TODO GET FROM ENVIROMENTS VARIABLES
	username := "root"    //TODO GET FROM ENVIROMENTS VARIABLES
	password := "example" //TODO GET FROM ENVIROMENTS VARIABLES

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

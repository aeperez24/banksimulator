package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USER_DATABASE_NAME = "promotions"
const USER_COLLECTION = "products"

type UserRepository interface {
	FindUserByName(username string) User
}

type userRepositoryMongoRepository struct {
	DBClient *mongo.Client
}

func (repo userRepositoryMongoRepository) FindUserByName(username string) User {

	var user User
	filter := bson.D{primitive.E{Key: "username", Value: username}}

	collection := repo.DBClient.Database(USER_DATABASE_NAME).Collection(USER_COLLECTION)
	collection.FindOne(context.TODO(), filter).Decode(&user)
	return user
}

func NewUserMongoRepository(DBClient *mongo.Client) UserRepository {
	return userRepositoryMongoRepository{DBClient: DBClient}

}

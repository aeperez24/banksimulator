package model

import (
	"aeperez24/banksimulator/config"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USER_DATABASE_NAME = "bank"
const USER_COLLECTION = "user"

type UserRepository interface {
	FindUserByName(username string) User
	CreateUser(user User) (interface{}, error)
	FindUserByIdDocument(idDocument string) User
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

func (repo userRepositoryMongoRepository) FindUserByIdDocument(idDocument string) User {

	var user User
	filter := bson.D{primitive.E{Key: "iddocument", Value: idDocument}}

	collection := repo.DBClient.Database(USER_DATABASE_NAME).Collection(USER_COLLECTION)
	collection.FindOne(context.TODO(), filter).Decode(&user)
	return user
}
func (repo userRepositoryMongoRepository) CreateUser(user User) (interface{}, error) {

	collection := repo.DBClient.Database(USER_DATABASE_NAME).Collection(USER_COLLECTION)
	resultID1, error := collection.InsertOne(context.TODO(), user)
	return resultID1.InsertedID, error
}

func NewUserMongoRepository(DBClient config.MongoCofig) UserRepository {
	return userRepositoryMongoRepository{DBClient: DBClient.DB}

}

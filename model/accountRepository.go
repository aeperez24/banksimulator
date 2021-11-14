package model

import (
	"aeperez24/banksimulator/config"
	"context"

	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ACCOUNT_COLLECTION = "account"

type AccountRepository interface {
	FindAccountByAccountNumber(account string) Account
	ModifyBalanceForAccount(accountNumber string, amount float32) bool
}

type accountMongoRepository struct {
	dbClient     *mongo.Client
	databaseName string
}

func (repo accountMongoRepository) FindAccountByAccountNumber(accountNumber string) Account {
	var account Account
	filter := bson.D{primitive.E{Key: "AccountNumber", Value: accountNumber}}
	collection := repo.dbClient.Database(repo.databaseName).Collection(ACCOUNT_COLLECTION)
	collection.FindOne(context.TODO(), filter).Decode(&account)
	return account
}
func (repo accountMongoRepository) ModifyBalanceForAccount(accountNumber string, amount float32) bool {
	filter := bson.D{primitive.E{Key: "AccountNumber", Value: accountNumber}}
	collection := repo.dbClient.Database(repo.databaseName).Collection(ACCOUNT_COLLECTION)
	update := bson.D{{"$inc", bson.D{
		{"Balance", amount},
	}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return err == nil
}
func NewAccountMongoRepository(DBConfig config.MongoCofig) AccountRepository {

	return accountMongoRepository{dbClient: DBConfig.DB, databaseName: DBConfig.DatabaseName}
}

package integrationtest

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/model"
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestSaveUser(t *testing.T) {
	RunTestWithIntegrationServerGin(func(port string) {
		username := "username_for_testing"
		dbConfig := config.BuildDBConfig()
		user := model.User{Active: true, Username: username, Password: "pass", IDDocument: "document"}
		repo := model.NewUserMongoRepository(dbConfig)
		repo.CreateUser(user)
		collection := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.USER_COLLECTION)
		defer collection.DeleteOne(context.TODO(), bson.M{"username": username})
		result := repo.FindUserByName(username)
		if result != user {
			t.Errorf("expected %v and received %v", user, result)
		}
	})
}

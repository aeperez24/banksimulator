package integrationtest

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/handler"
	"aeperez24/banksimulator/model"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestSaveTransaction(t *testing.T) {
	DBConfig := config.BuildDBConfig()
	repo := model.NewAccountMongoRepository(DBConfig)
	ids := createAccountForTests(DBConfig)
	createAccountForTests(DBConfig)
	defer deleteAccountForTests(DBConfig, ids)
	repo.SaveTransaction("12", model.Transaction{AccountFrom: "account1Number", AccountTo: "account2Number", Amount: 10})
}

func TestGetBalance(t *testing.T) {
	port := "11080"
	DBConfig := config.BuildDBConfig()
	repo := model.NewAccountMongoRepository(DBConfig)
	achandler := handler.NewAccountHandler(repo)
	server := handler.NewServer(":"+port, achandler)
	ids := createAccountForTests(DBConfig)
	defer deleteAccountForTests(DBConfig, ids)
	go server.Start()
	defer server.Stop()

	api := fmt.Sprintf("http://localhost:%s/balance/account1Number", port)

	resp, _ := http.Get(api)
	body, _ := ioutil.ReadAll(resp.Body)
	//TODO ASSERT
	println(string(body))

}

func TestGetTrnsactions(t *testing.T) {

}

func createAccountForTests(dbConfig config.MongoCofig) []interface{} {
	collection := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.ACCOUNT_COLLECTION)

	account1 := model.Account{
		AccountNumber: "account1Number",
		Balance:       100,
	}
	account2 := model.Account{
		AccountNumber: "account2Number",
		Balance:       100,
	}
	resultID1, error1 := collection.InsertOne(context.TODO(), account1)
	resultID2, error2 := collection.InsertOne(context.TODO(), account2)
	if error1 != nil {
		println(error1)
		panic(error1)
	}
	if error2 != nil {
		println(error1)
		panic(error1)
	}
	result := []interface{}{
		resultID1.InsertedID, resultID2.InsertedID,
	}
	return result

}

func deleteAccountForTests(dbConfig config.MongoCofig, idaToDelte []interface{}) {
	collection := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.ACCOUNT_COLLECTION)
	for _, id := range idaToDelte {
		collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	}

}

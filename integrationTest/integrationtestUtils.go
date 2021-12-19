package integrationtest

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/handler"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func RunTestWithIntegrationServer(testFunc func(port string)) {
	DBConfig := config.BuildDBConfig()
	server, port := createTestServer(DBConfig)
	ids := createAccountForTests(DBConfig)
	defer deleteAccountForTests(DBConfig, ids)
	go server.Start()
	defer server.Stop()
	testFunc(port)
}
func createTestServer(DBConfig config.MongoCofig) (port.Server, string) {
	port := "11080"
	repo := model.NewAccountMongoRepository(DBConfig)
	achandler := handler.NewAccountHandler(repo)
	return handler.NewServer(":"+port, achandler), port
}

func createAccountForTests(dbConfig config.MongoCofig) []interface{} {
	collection := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.ACCOUNT_COLLECTION)

	account1 := model.Account{
		AccountNumber: "account1Number",
		Balance:       100,
		Transactions: []model.Transaction{{
			AccountTo: "account1Number",
			Amount:    100,
			Type:      model.DepositType,
		}},
	}
	account2 := model.Account{
		AccountNumber: "account2Number",
		Balance:       100,
		Transactions: []model.Transaction{{
			AccountTo: "account1Number",
			Amount:    50,
			Type:      model.DepositType,
		}, {
			AccountTo: "account1Number",
			Amount:    50,
			Type:      model.DepositType,
		}},
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

package integrationtest

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/handler"
	"aeperez24/banksimulator/middleware"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/port"
	"aeperez24/banksimulator/services"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func RunTestWithIntegrationServer(testFunc func(port string)) {
	DBConfig := config.BuildDBConfig()
	server, port := createTestServer(DBConfig)
	idAccounts := createAccountForTests(DBConfig)
	defer deleteAccountForTests(DBConfig, idAccounts)
	idUsers := createUserForTest(DBConfig)
	defer deleteUsersForTests(DBConfig, idUsers)
	go server.Start()
	defer server.Stop()
	testFunc(port)
}
func createTestServer(DBConfig config.MongoCofig) (port.Server, string) {
	port := "11080"
	accountRepo := model.NewAccountMongoRepository(DBConfig)
	userRepo := model.NewUserMongoRepository(DBConfig)
	userService := services.NewUserService(userRepo)
	achandler := handler.NewAccountHandler(accountRepo)
	userHandler := handler.NewUserhandler(accountRepo, userService)
	tokenService := services.NewTokenService("testKey")
	authMiddleware := middleware.NewAuthenticationMiddlware(tokenService)
	authHandler := handler.NewAuthenticationHandler(userService, tokenService)

	config := handler.HandlerConfig{
		AccountHandler:        achandler,
		UserHandler:           userHandler,
		AuthenticationHandler: authHandler,
	}
	serverConfig := handler.ServerConfiguration{
		Port:             ":" + port,
		MiddleWareConfig: middleware.MiddlewareConfig{AuthenticationMiddleware: authMiddleware},
		HandlerConfig:    config,
	}
	return handler.NewServer(serverConfig), port
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

func createUserForTest(dbConfig config.MongoCofig) []interface{} {

	user1 := model.User{Username: "user1", Password: "pass1", IDDocument: "account1Number"}
	user2 := model.User{Username: "user2", Password: "pass2", IDDocument: "account2Number"}
	collection := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.USER_COLLECTION)
	resultID1, error1 := collection.InsertOne(context.TODO(), user1)
	resultID2, error2 := collection.InsertOne(context.TODO(), user2)
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

func deleteUsersForTests(dbConfig config.MongoCofig, idaToDelte []interface{}) {
	collection := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.USER_COLLECTION)
	for _, id := range idaToDelte {
		collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	}

}

func GetJWTTokenForUser1() string {
	tokenService := services.NewTokenService("testKey")
	res, _ := tokenService.CreateToken(dto.BasicUserDto{
		Username:   "user1",
		IDDocument: "account1Number",
	})
	return res
}

func GetJWTTokenForUser2() string {
	tokenService := services.NewTokenService("testKey")
	res, _ := tokenService.CreateToken(dto.BasicUserDto{
		Username:   "user2",
		IDDocument: "account2Number",
	})
	return res
}

func ExecuteHttpPostCall(url string, bodyInterface interface{}, headers map[string]string) ([]byte, *http.Response, error) {
	body, _ := json.Marshal(bodyInterface)
	postBuffer := bytes.NewBuffer(body)

	req, _ := http.NewRequest("POST", url, postBuffer)
	for name, value := range headers {
		req.Header.Add(name, value)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, _ := client.Do((req))
	bodyresp, err := ioutil.ReadAll(resp.Body)
	return bodyresp, resp, err
}

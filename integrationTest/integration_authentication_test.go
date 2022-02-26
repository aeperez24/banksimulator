package integrationtest

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TestAuthUser(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {
		username := "username_for_integration_testing"
		document := "document_for_integration_testing"
		password := "pass"
		dbConfig := config.BuildDBConfig()
		collectionUser := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.USER_COLLECTION)
		defer collectionUser.DeleteOne(context.TODO(), bson.M{"username": username})

		collectionAccount := dbConfig.DB.Database(dbConfig.DatabaseName).Collection(model.ACCOUNT_COLLECTION)

		defer collectionAccount.DeleteOne(context.TODO(), bson.M{"accountnumber": document})

		userdto := dto.UserWithPasswordDto{
			BasicUserDto: dto.BasicUserDto{
				Username:   username,
				IDDocument: document,
			},
			Password: password,
		}
		body, _ := json.Marshal(userdto)
		postBufferSignUp := bytes.NewBuffer(body)

		postBufferSignIn := bytes.NewBuffer(body)

		apiSignUp := fmt.Sprintf("http://localhost:%s/user/signup", port)
		apiSignIn := fmt.Sprintf("http://localhost:%s/user/signin", port)
		reqSignUp, _ := http.NewRequest("POST", apiSignUp, postBufferSignUp)
		reqSignIn, _ := http.NewRequest("POST", apiSignIn, postBufferSignIn)

		client := &http.Client{
			Timeout: time.Second * 10,
		}
		resp1, _ := client.Do((reqSignUp))
		bresp1, _ := ioutil.ReadAll(resp1.Body)
		println(string(bresp1))

		resp, _ := client.Do((reqSignIn))
		bodyresp, _ := ioutil.ReadAll(resp.Body)
		//TODO ASSERT
		println(string(bodyresp))

	})
}

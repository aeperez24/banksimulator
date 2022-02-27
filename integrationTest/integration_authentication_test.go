package integrationtest

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/test"
	"context"
	"fmt"
	"testing"

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
		apiSignUp := fmt.Sprintf("http://localhost:%s/user/signup", port)
		apiSignIn := fmt.Sprintf("http://localhost:%s/user/signin", port)

		ExecuteHttpPostCall(apiSignUp, userdto, nil)
		bodyresp, resp, err := ExecuteHttpPostCall(apiSignIn, userdto, nil)
		assertHelper := test.AssertHelper{
			T: t,
		}
		assertHelper.Assert(200, resp.StatusCode)
		assertHelper.Assert(nil, err)
		println(string(bodyresp))

	})
}

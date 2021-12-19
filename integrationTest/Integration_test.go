package integrationtest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

/* func TestSaveTransaction(t *testing.T) {
	DBConfig := config.BuildDBConfig()
	repo := model.NewAccountMongoRepository(DBConfig)
	ids := createAccountForTests(DBConfig)
	createAccountForTests(DBConfig)
	defer deleteAccountForTests(DBConfig, ids)
	repo.SaveTransaction("12", model.Transaction{AccountFrom: "account1Number", AccountTo: "account2Number", Amount: 10})
} */

func TestGetBalance(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {
		api := fmt.Sprintf("http://localhost:%s/balance/account1Number", port)
		resp, _ := http.Get(api)
		body, _ := ioutil.ReadAll(resp.Body)
		//TODO ASSERT
		println(string(body))

	})
}
func TestGetTrnsactions(t *testing.T) {

}

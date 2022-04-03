package integrationtest

import (
	"aeperez24/banksimulator/dto"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {
		api := fmt.Sprintf("http://localhost:%s/account/balance/account1Number", port)
		req, _ := http.NewRequest("GET", api, nil)
		token := GetJWTTokenForUser1()
		req.Header.Add("Authorization", fmt.Sprintf("bearer %v", token))
		client := &http.Client{
			Timeout: time.Second * 10,
		}
		resp, _ := client.Do((req))
		body, _ := ioutil.ReadAll(resp.Body)
		//TODO ASSERT
		println(string(body))
		println(resp.StatusCode)
		assert.Equal(t, "{\"Data\":100}", string(body), "Error get balance")

	})
}
func TestGetTransactions(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {
		api := fmt.Sprintf("http://localhost:%s/transaction/account1Number", port)
		req, _ := http.NewRequest("GET", api, nil)
		token := GetJWTTokenForUser1()
		req.Header.Add("Authorization", fmt.Sprintf("bearer %v", token))
		client := &http.Client{
			Timeout: time.Second * 10,
		}
		resp, _ := client.Do((req))

		body, _ := ioutil.ReadAll(resp.Body)

		expected := "\"AccountTo\":\"account1Number\",\"Amount\":100,\"Date\":\"0001-01-01T00:00:00Z\",\"Type\":\"Deposit\"}"
		println(string(body))
		assert.Contains(t, string(body), expected, "Error get transactions")

	})

}

func TestTransferMoney(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {

		transaction := dto.TransferRequest{
			FromAccount: "account1Number",
			ToAccount:   "account2Number",
			Amount:      20,
		}

		api := fmt.Sprintf("http://localhost:%s/account/transfer/", port)
		token := GetJWTTokenForUser1()
		headers := make(map[string]string)
		headers["Authorization"] = fmt.Sprintf("bearer %v", token)
		bodyresp, _, _ := ExecuteHttpPostCall(api, transaction, headers)

		//TODO  fix test
		println(string(bodyresp))
		//		assert.Equal(t, "{\"Data\":80}", string(bodyresp), "Error transfer")

	})

}

func TestDepositMoney(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {

		transaction := dto.DepositRequest{
			ToAccount: "account1Number",
			Amount:    20,
		}

		api := fmt.Sprintf("http://localhost:%s/account/deposit/", port)
		headers := make(map[string]string)
		token := GetJWTTokenForUser1()
		headers["Authorization"] = fmt.Sprintf("bearer %v", token)
		bodyresp, _, _ := ExecuteHttpPostCall(api, transaction, headers)

		//TODO ASSERT
		println(string(bodyresp))

	})

}

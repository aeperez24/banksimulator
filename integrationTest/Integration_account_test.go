package integrationtest

import (
	"aeperez24/banksimulator/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetBalance(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {
		api := fmt.Sprintf("http://localhost:%s/balance/account1Number", port)
		resp, _ := http.Get(api)
		body, _ := ioutil.ReadAll(resp.Body)
		//TODO ASSERT
		println(string(body))

	})
}
func TestGetTransactions(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {
		api := fmt.Sprintf("http://localhost:%s/transaction/account1Number", port)
		resp, _ := http.Get(api)
		body, _ := ioutil.ReadAll(resp.Body)
		//TODO ASSERT
		println(string(body))

	})

}

func TestTransferMoney(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {

		transaction := dto.TransferRequest{
			FromAccount: "account1Number",
			ToAccount:   "account2Number",
			Amount:      20,
		}

		body, _ := json.Marshal(transaction)
		postBuffer := bytes.NewBuffer(body)
		api := fmt.Sprintf("http://localhost:%s/transfer/", port)
		resp, _ := http.Post(api, "application/json", postBuffer)
		bodyresp, _ := ioutil.ReadAll(resp.Body)
		//TODO ASSERT
		println(string(bodyresp))

	})

}

func TestDepositMoney(t *testing.T) {
	RunTestWithIntegrationServer(func(port string) {

		transaction := dto.DepositRequest{
			ToAccount: "account1Number",
			Amount:    20,
		}

		body, _ := json.Marshal(transaction)
		postBuffer := bytes.NewBuffer(body)
		api := fmt.Sprintf("http://localhost:%s/deposit/", port)
		resp, _ := http.Post(api, "application/json", postBuffer)
		bodyresp, _ := ioutil.ReadAll(resp.Body)
		//TODO ASSERT
		println(string(bodyresp))

	})

}

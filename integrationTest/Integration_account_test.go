package integrationtest

import (
	"aeperez24/banksimulator/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
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
		api := fmt.Sprintf("http://localhost:%s/account/transfer/", port)
		req, _ := http.NewRequest("POST", api, postBuffer)
		token := GetJWTTokenForUser1()
		req.Header.Add("Authorization", fmt.Sprintf("bearer %v", token))

		client := &http.Client{
			Timeout: time.Second * 10,
		}
		resp, _ := client.Do((req))
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

		api := fmt.Sprintf("http://localhost:%s/account/deposit/", port)
		req, _ := http.NewRequest("POST", api, postBuffer)
		token := GetJWTTokenForUser1()
		client := &http.Client{
			Timeout: time.Second * 10,
		}
		req.Header.Add("Authorization", fmt.Sprintf("bearer %v", token))
		resp, _ := client.Do((req))
		bodyresp, _ := ioutil.ReadAll(resp.Body)

		//TODO ASSERT
		println(string(bodyresp))

	})

}

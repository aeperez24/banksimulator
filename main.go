package main

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/model"
	"aeperez24/banksimulator/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/balance/{AccountNumber}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		service := getAccountService(vars["AccountNumber"])
		respondWithJSON(w, 200, service.GetBalance())
	})

	r.HandleFunc("/transfer/", func(w http.ResponseWriter, r *http.Request) {
		var transferRequest dto.TransferRequest

		err := json.NewDecoder(r.Body).Decode(&transferRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		service := getAccountService(transferRequest.FromAccount)
		service.TransferMoneyTo(transferRequest.ToAccount, transferRequest.Amount)
		respondWithJSON(w, 200, service.GetBalance())

	})

	r.HandleFunc("/deposit/", func(w http.ResponseWriter, r *http.Request) {
		var depositRequest dto.DepositRequest

		err := json.NewDecoder(r.Body).Decode(&depositRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		service := getAccountService(depositRequest.ToAccount)
		service.Deposit(depositRequest.Amount)
		respondWithJSON(w, 200, service.GetBalance())

	})

	http.ListenAndServe(":8080", r)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getAccountService(accountNumber string) services.AccountService {
	repo := model.NewAccountMongoRepository(config.DBConfig)
	return services.NewAccountService(accountNumber, repo)

}

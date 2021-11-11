package main

import "fmt"
import "net/http"

import "github.com/gorilla/mux"
import "aeperez24/banksimulator/config"
import "aeperez24/banksimulator/services"
import "aeperez24/banksimulator/model"
func main() {
    r := mux.NewRouter()

    r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        fmt.Print("api called")
		service := getAccountService("1234")
		service.Deposit(100)
		respondWithJSON(w,200,nil)
    })

    http.ListenAndServe(":80", r)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
   // response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
   // w.Write(response)
}

func getAccountService(accountNumber string) services.AccountService{
	 repo := model.NewAccountMongoRepository(config.DB)
	 return services.NewAccountService(accountNumber,repo)

}
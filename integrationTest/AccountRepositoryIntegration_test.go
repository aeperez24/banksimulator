package integrationtest

import (
	"aeperez24/banksimulator/config"
	"aeperez24/banksimulator/model"
	"testing"
)

func TestSaveTransaction(t *testing.T) {
	DBConfig := config.BuildDBConfig()
	repo := model.NewAccountMongoRepository(DBConfig)
	repo.SaveTransaction("12", model.Transaction{AccountFrom: "12", AccountTo: "1234", Amount: 400})

}

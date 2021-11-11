package model
import "go.mongodb.org/mongo-driver/mongo"
type AccountRepository interface {
	FindAccountByAccountNumber(account string) Account
	ModifyBalanceForAccount(accountNumber string, amount float32) bool
}

type accountMongoRepository struct{
	DBClient *mongo.Client
}

func (r accountMongoRepository)FindAccountByAccountNumber(account string)Account{
	return Account{}
}
func (r accountMongoRepository)ModifyBalanceForAccount(accountNumber string, amount float32)bool{
	return false
}
func NewAccountMongoRepository(DBClient *mongo.Client)AccountRepository{

	return accountMongoRepository{DBClient: DBClient}
}
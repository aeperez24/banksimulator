package model

type User struct {
	Username          string
	Password          string
	active            bool
	AccountNumberList []string
}

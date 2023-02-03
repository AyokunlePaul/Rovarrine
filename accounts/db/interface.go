package db

type Database interface {
	CreateAccount(Customer)
	GetUserInformation(string) *Customer
}

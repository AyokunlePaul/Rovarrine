package db

import "Rovarrine/accounts/rpc"

type Database interface {
	CreateAccount(Customer)
	GetUserInformation(string) *rpc.GetUserInformationResponse
}

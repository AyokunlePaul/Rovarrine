package db

import (
	"Rovarrine/accounts/rpc"
	"fmt"
	"github.com/hashicorp/go-memdb"
	"log"
)

type inMem struct {
	Database *memdb.MemDB
}

type Customer struct {
	CustomerId string
	FirstName  string
	LastName   string
}

func NewInMemoryDatabase() Database {
	schema := &memdb.DBSchema{Tables: map[string]*memdb.TableSchema{
		"users": {
			Name: "users",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "customer_id"},
				},
			},
		},
	}}
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(fmt.Sprintf("error creating database: %v", err))
	}
	return &inMem{Database: db}
}

func (i *inMem) CreateAccount(customer Customer) {
	txn := i.Database.Txn(true)
	err := txn.Insert("users", customer)
	if err != nil {
		log.Fatalf("error saving data: %v\n", customer)
	}
	txn.Commit()
}

func (i *inMem) GetUserInformation(userId string) *rpc.GetUserInformationResponse {
	return nil
}

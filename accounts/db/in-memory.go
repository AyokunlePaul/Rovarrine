package db

import (
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
					Indexer: &memdb.StringFieldIndex{Field: "CustomerId"},
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
		log.Fatalf("error saving data: %v\n", err)
	}
	txn.Commit()
}

func (i *inMem) GetUserInformation(userId string) *Customer {
	txn := i.Database.Txn(false)
	defer txn.Abort()

	scan, err := txn.First("users", "CustomerId", userId)
	if err != nil {
		return nil
	}

	return scan.(*Customer)
}

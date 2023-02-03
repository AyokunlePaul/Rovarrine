package db

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
)

type inMem struct {
	Database *memdb.MemDB
}

func NewInMemoryDatabase() Database {
	schema := &memdb.DBSchema{Tables: map[string]*memdb.TableSchema{
		"transactions": {
			Name: "transactions",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "TransactionId"},
				},
				"user_id": {
					Name:    "user_id",
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

func (i *inMem) Create(transaction Transaction) {
	txn := i.Database.Txn(true)
	err := txn.Insert("transactions", transaction)
	if err != nil {
		fmt.Printf("error saving transaction data: %v", err)
		return
	}
}

func (i *inMem) Get(customerId string) (transactions []Transaction, err error) {
	txn := i.Database.Txn(false)
	defer txn.Abort()

	scan, err := txn.Get("transactions", "user_id", customerId)
	if err != nil {
		return
	}
	for {
		if currTxn := scan.Next(); currTxn != nil {
			transactions = append(transactions, currTxn.(Transaction))
			break
		}
	}

	return
}

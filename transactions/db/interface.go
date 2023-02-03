package db

type Database interface {
	Create(Transaction)
	Get(string) ([]Transaction, error)
}

type Transaction struct {
	TransactionId string
	CustomerId    string
	Balance       int64
}

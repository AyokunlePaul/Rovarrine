package rpc

import (
	"Rovarrine/transactions/db"
	"context"
	"github.com/google/uuid"
	"net/http"
)

type Server struct {
	Db   db.Database
	Port string
}

func InitializeServer() *Server {
	cache := db.NewInMemoryDatabase()
	return &Server{
		Db:   cache,
		Port: config().Port,
	}
}

func (s *Server) RecordTransaction(_ context.Context, req *RecordTransactionRequest) (*RecordTransactionResponse, error) {
	txnId, _ := uuid.NewUUID()
	data := db.Transaction{
		TransactionId: txnId.String(),
		CustomerId:    req.UserId,
		Balance:       int64(req.Amount),
	}
	s.Db.Create(data)
	return &RecordTransactionResponse{
		Status:  http.StatusOK,
		Success: true,
		Message: "transaction saved successfully",
	}, nil
}

func (s *Server) GetTransactions(_ context.Context, req *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	resp, err := s.Db.Get(req.UserId)
	if err != nil {
		return nil, err
	}
	var txns []*Transaction

	for _, txn := range resp {
		txns = append(txns, ToRpcTransaction(txn))
	}

	return &GetTransactionsResponse{
		Amount:       float64(calculateBalance(resp)),
		Transactions: txns,
		Date:         nil,
	}, nil
}

func ToRpcTransaction(txn db.Transaction) *Transaction {
	return &Transaction{
		Id:     txn.TransactionId,
		Amount: txn.Balance,
	}
}

func calculateBalance(txns []db.Transaction) int64 {
	amount := int64(0)
	for _, txn := range txns {
		amount += txn.Balance
	}
	return amount
}

func (s *Server) mustEmbedUnimplementedTransactionServiceServer() {}

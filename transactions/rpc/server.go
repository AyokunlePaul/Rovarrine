package rpc

import (
	"Rovarrine/transactions"
	"Rovarrine/transactions/db"
	"context"
	"github.com/google/uuid"
	"net/http"
)

type Server struct {
	Db   db.Database
	Port string
}

func InitializeServer(config *transactions.Config) *Server {
	cache := db.NewInMemoryDatabase()
	return &Server{
		Db:   cache,
		Port: config.Port,
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

func (s *Server) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, nil
}

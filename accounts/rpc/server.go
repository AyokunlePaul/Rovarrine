package rpc

import (
	"Rovarrine/accounts/conf"
	"Rovarrine/accounts/db"
	txnRpc "Rovarrine/transactions/rpc"
	"context"
	"errors"
	"fmt"
	"github.com/goombaio/namegenerator"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Db        db.Database
	Port      string
	TxnClient txnRpc.TransactionServiceClient
}

func InitializeServer(config *conf.AccountConfig, TxnClient txnRpc.TransactionServiceClient) *Server {
	cache := db.NewInMemoryDatabase()
	return &Server{
		Db:        cache,
		Port:      config.ServerPort,
		TxnClient: TxnClient,
	}
}

func (s *Server) CreateAccount(ctx context.Context, account *Account) (*CreateAccountResponse, error) {
	seed := time.Now().UTC().UnixNano()
	gen := namegenerator.NewNameGenerator(seed)

	names := strings.Split(gen.Generate(), "-")
	firstName := cases.Title(language.Make("en_US")).String(names[0])
	lastName := cases.Title(language.Make("en_US")).String(names[1])

	data := db.Customer{
		CustomerId: account.CustomerId,
		FirstName:  firstName,
		LastName:   lastName,
	}
	if account.InitialCredit >= 0.0 {
		_, err := s.TxnClient.RecordTransaction(ctx, &txnRpc.RecordTransactionRequest{
			Amount: account.InitialCredit,
			UserId: account.CustomerId,
			Date:   nil,
		})
		if err != nil {
			fmt.Printf("unable to save transaction --> error: %v", err)
			return nil, err
		}
	}
	s.Db.CreateAccount(data)

	return &CreateAccountResponse{
		Status:  http.StatusCreated,
		Success: true,
		Message: "customer account created successfully",
	}, nil
}

func (s *Server) GetUserInformation(ctx context.Context, req *UserInformationRequest) (*GetUserInformationResponse, error) {
	txns, err := s.TxnClient.GetTransactions(ctx, &txnRpc.GetTransactionsRequest{UserId: req.UserId})
	if err != nil {
		fmt.Printf("unable to fetch transactions --> error: %v", err)
	}
	customer := s.Db.GetUserInformation(req.UserId)
	if customer == nil {
		return nil, errors.New("user doesn't exist")
	}
	return &GetUserInformationResponse{
		Surname:      customer.LastName,
		Name:         customer.FirstName,
		Balance:      fmt.Sprintf("%f", txns.Amount),
		Transactions: nil,
	}, nil
}

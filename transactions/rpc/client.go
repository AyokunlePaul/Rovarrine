package rpc

import (
	"Rovarrine/transactions"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func config() *transactions.Config {
	return &transactions.Config{
		AppName: "transaction-server",
		Port:    ":22002",
		SvcUrl:  "localhost:22002",
	}
}

func InitTransactionClient() TransactionServiceClient {
	cc, err := grpc.Dial(config().SvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("could not connect to account service --> err: %v", err)
		return nil
	}

	return NewTransactionServiceClient(cc)
}

package rpc

import (
	"Rovarrine/accounts/conf"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAccountService(config *conf.AccountConfig) AccountServiceClient {
	cc, err := grpc.Dial(config.SvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("could not connect to account service --> err: %v", err)
		return nil
	}

	return NewAccountServiceClient(cc)
}

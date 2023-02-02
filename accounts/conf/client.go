package conf

import (
	"Rovarrine/accounts/rpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client rpc.AccountServiceClient
}

func InitAccountService(config *AccountConfig) rpc.AccountServiceClient {
	cc, err := grpc.Dial(config.SvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("could not connect to account service --> err: %v", err)
		return nil
	}

	return rpc.NewAccountServiceClient(cc)
}

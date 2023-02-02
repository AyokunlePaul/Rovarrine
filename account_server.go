package main

import (
	"Rovarrine/accounts/conf"
	"Rovarrine/accounts/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Account server
	server := conf.InitializeServer(&conf.AccountConfig{
		AppName:    "account-server",
		ServerPort: ":12290",
	})
	lis, err := net.Listen("tcp", server.Port)
	if err != nil {
		log.Fatalln("error listening on tcp address")
	}

	grpcServer := grpc.NewServer()
	rpc.RegisterAccountServiceServer(grpcServer, server)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalln("error serving to listener")
	}
}

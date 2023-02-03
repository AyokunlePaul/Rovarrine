package main

import (
	txnRpc "Rovarrine/transactions/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Account server
	server := txnRpc.InitializeServer()
	lis, err := net.Listen("tcp", server.Port)
	if err != nil {
		log.Fatalln("error listening on tcp address")
	}

	grpcServer := grpc.NewServer()
	txnRpc.RegisterTransactionServiceServer(grpcServer, server)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalln("error serving to listener")
	}
}

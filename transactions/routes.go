package transactions

import (
	"Rovarrine/transactions/rpc"
	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	CreateCurrentAccount(*gin.Context)
	GetUserInformation(*gin.Context)
}

type handler struct {
	Client rpc.TransactionServiceClient
}

func NewHandler(Client rpc.TransactionServiceClient) TransactionHandler {
	return &handler{
		Client: Client,
	}
}

func (h *handler) CreateCurrentAccount(ctx *gin.Context) {

}

func (h *handler) GetUserInformation(ctx *gin.Context) {

}

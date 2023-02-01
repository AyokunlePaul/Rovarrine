package accounts

import (
	"Rovarrine/accounts/rpc"
	"github.com/gin-gonic/gin"
)

type AccountHandler interface {
	CreateCurrentAccount(*gin.Context)
	GetUserInformation(*gin.Context)
}

type handler struct {
	Client rpc.AccountServiceClient
}

func NewHandler(Client rpc.AccountServiceClient) AccountHandler {
	return &handler{
		Client: Client,
	}
}

func (h *handler) CreateCurrentAccount(ctx *gin.Context) {

}

func (h *handler) GetUserInformation(ctx *gin.Context) {

}

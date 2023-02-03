package accounts

import (
	"Rovarrine/accounts/conf"
	"Rovarrine/accounts/rpc"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, config *conf.AccountConfig) {
	client := rpc.InitAccountService(config)
	handler := NewHandler(client)

	r.POST("/create-account", handler.CreateCurrentAccount)
	r.GET("/user/:user_id", handler.GetUserInformation)
}

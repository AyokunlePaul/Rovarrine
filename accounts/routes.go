package accounts

import (
	"Rovarrine/accounts/conf"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, config *conf.AccountConfig) {
	client := conf.InitAccountService(config)
	handler := NewHandler(client)

	r.POST("/create-account", handler.CreateCurrentAccount)
	r.POST("/user/:user_id", handler.GetUserInformation)
}

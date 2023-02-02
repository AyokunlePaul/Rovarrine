package main

import (
	"Rovarrine/accounts"
	"Rovarrine/accounts/conf"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Account client
	r := gin.Default()
	accounts.RegisterRoutes(r, &conf.AccountConfig{
		AppName:    "account-app",
		SvcUrl:     "localhost:12290",
		ServerPort: ":12290",
	})

	err := r.Run(":1122")
	if err != nil {
		panic(fmt.Sprintf("unable to start app --> err: %v", err))
	}
}

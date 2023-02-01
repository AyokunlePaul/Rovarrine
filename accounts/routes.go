package accounts

import (
	"Rovarrine/accounts/rpc"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	req := Request{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, struct {
			Success bool
			Message string
		}{
			Success: false,
			Message: "invalid request data",
		})
		return
	}

	err = req.IsValidRequestData()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, struct {
			Success bool
			Message string
		}{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	customerId := req.Data["customer_id"].(string)
	// Initial credit should be sent as kobo
	initialCredit := req.Data["initial_credit"].(int64)

	resp, err := h.Client.CreateAccount(context.Background(), &rpc.Account{
		CustomerId:    customerId,
		InitialCredit: float64(initialCredit / 100),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, struct {
			Success bool
			Message string
		}{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, struct {
		Success bool
		Message string
		Data    interface{}
	}{
		Success: true,
		Message: "account created",
		Data:    resp,
	})
}

func (h *handler) GetUserInformation(ctx *gin.Context) {
	userId := ctx.Param("userId")
	if strings.TrimSpace(userId) == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, struct {
			Success bool
			Message string
		}{
			Success: false,
			Message: "user id is required",
		})
		return
	}
	resp, err := h.Client.GetUserInformation(context.Background(), &rpc.UserInformationRequest{UserId: userId})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, struct {
			Success bool
			Message string
		}{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, struct {
		Success bool
		Message string
		Data    interface{}
	}{
		Success: true,
		Message: "user information fetched",
		Data:    resp,
	})
}

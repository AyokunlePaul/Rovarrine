package conf

import (
	"Rovarrine/accounts/db"
	"Rovarrine/accounts/rpc"
	"context"
	"github.com/goombaio/namegenerator"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Db   db.Database
	Port string
}

func InitializeServer(config *AccountConfig) *Server {
	cache := db.NewInMemoryDatabase()
	return &Server{
		Db:   cache,
		Port: config.ServerPort,
	}
}

func (s *Server) CreateAccount(ctx context.Context, account *rpc.Account) (*rpc.CreateAccountResponse, error) {
	seed := time.Now().UTC().UnixNano()
	gen := namegenerator.NewNameGenerator(seed)

	names := strings.Split(gen.Generate(), "-")
	firstName := cases.Title(language.Make("en_US")).String(names[0])
	lastName := cases.Title(language.Make("en_US")).String(names[1])

	data := db.Customer{
		CustomerId: account.CustomerId,
		FirstName:  firstName,
		LastName:   lastName,
	}

	s.Db.CreateAccount(data)

	return &rpc.CreateAccountResponse{
		Status:  http.StatusCreated,
		Success: true,
		Message: "customer account created successfully",
	}, nil
}

func (s *Server) GetUserInformation(context.Context, *rpc.UserInformationRequest) (*rpc.GetUserInformationResponse, error) {
	return nil, nil
}

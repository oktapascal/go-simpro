package model

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/helper"
	"net/http"
)

type (
	Client struct {
		ID      string
		Name    string
		Address string
		Phone   string
		Status  string
	}

	ClientPIC struct {
		ID       int
		IDClient string
		Name     string
		Phone    string
		Email    string
		Address  string
	}

	SaveRequestClientPIC struct {
		Name    string `json:"name" validate:"required,min=1,max=50"`
		Email   string `json:"email" validate:"required,email,min=1,max=50"`
		Phone   string `json:"phone" validate:"required,min=1,max=13"`
		Address string `json:"address" validate:"required,min=1,max=100"`
	}

	SaveRequestClient struct {
		Name      string                 `json:"name" validate:"required,min=1,max=50"`
		Address   string                 `json:"address" validate:"required,min=1,max=100"`
		Phone     string                 `json:"phone" validate:"required,min=11,max=13"`
		ClientPIC []SaveRequestClientPIC `json:"client_pic" validate:"required,minclientpic,dive"`
	}

	UpdateRequestClientPIC struct {
		ID       int    `json:"id"`
		IDClient string `json:"id_client"`
		Name     string `json:"name" validate:"required,min=1,max=50"`
		Email    string `json:"email" validate:"required,email,min=1,max=50"`
		Phone    string `json:"phone" validate:"required,min=1,max=13"`
		Address  string `json:"address" validate:"required,min=1,max=100"`
	}

	UpdateRequestClient struct {
		ID        string                   `json:"id" validate:"required"`
		Name      string                   `json:"name" validate:"required,min=1,max=50"`
		Address   string                   `json:"address" validate:"required,min=1,max=100"`
		Phone     string                   `json:"phone" validate:"required,min=11,max=13"`
		ClientPIC []UpdateRequestClientPIC `json:"client_pic" validate:"required,minclientpic,dive"`
	}

	ClientPICResponse struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Email   string `json:"email"`
		Address string `json:"address"`
	}

	ClientResponse struct {
		ID        string              `json:"id"`
		Name      string              `json:"name"`
		Address   string              `json:"address"`
		Phone     string              `json:"phone"`
		Status    string              `json:"status"`
		ClientPic []ClientPICResponse `json:"client_pic"`
	}

	ClientRepository interface {
		SaveClient(ctx context.Context, tx *sql.Tx, data *Client)
		SaveClientPIC(ctx context.Context, tx *sql.Tx, data *[]ClientPIC)
		UpdateClient(ctx context.Context, tx *sql.Tx, data *Client)
		UpdateClientPIC(ctx context.Context, tx *sql.Tx, data *[]ClientPIC)
		GetClients(ctx context.Context, tx *sql.Tx) *[]Client
		GetClientsWithPagination(ctx context.Context, tx *sql.Tx, params *helper.PaginationParams) *[]Client
		GetClient(ctx context.Context, tx *sql.Tx, id string) (*Client, error)
		GetClientPIC(ctx context.Context, tx *sql.Tx, id string) *[]ClientPIC
		DeleteClient(ctx context.Context, tx *sql.Tx, id string)
		DeleteClientPIC(ctx context.Context, tx *sql.Tx, IDClient string, id []int)
	}

	ClientService interface {
		SaveClient(ctx context.Context, request *SaveRequestClient) ClientResponse
		UpdateClient(ctx context.Context, request *UpdateRequestClient) ClientResponse
		GetClients(ctx context.Context) []ClientResponse
		GetClientsWithPagination(ctx context.Context, params *helper.PaginationParams) []ClientResponse
		GetClient(ctx context.Context, id string) ClientResponse
		DeleteClient(ctx context.Context, id string)
	}

	ClientHandler interface {
		SaveClient() http.HandlerFunc
		UpdateClient() http.HandlerFunc
		GetClients() http.HandlerFunc
		GetClientsWithPagination() http.HandlerFunc
		GetClient() http.HandlerFunc
		DeleteClient() http.HandlerFunc
	}
)

package model

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/helper"
	"net/http"
)

type (
	PIC struct {
		ID     string
		Name   string
		Email  string
		Phone  string
		Status string
	}

	SaveRequestPIC struct {
		Name  string `json:"name" validate:"required,min=1,max=50"`
		Email string `json:"email" validate:"required,email,min=1,max=50"`
		Phone string `json:"phone" validate:"required,min=1,max=13"`
	}

	UpdateRequestPIC struct {
		ID    string `json:"id"`
		Name  string `json:"name" validate:"required,min=1,max=50"`
		Email string `json:"email" validate:"required,email,min=1,max=50"`
		Phone string `json:"phone" validate:"required,min=1,max=13"`
	}

	PICResponse struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Phone  string `json:"phone"`
		Email  string `json:"email"`
		Status string `json:"status"`
	}

	PICRepository interface {
		SavePIC(ctx context.Context, tx *sql.Tx, data *PIC)
		UpdatePIC(ctx context.Context, tx *sql.Tx, data *PIC)
		GetPICs(ctx context.Context, tx *sql.Tx) *[]PIC
		GetPICsWithPagination(ctx context.Context, tx *sql.Tx, params *helper.PaginationParams) *[]PIC
		GetPIC(ctx context.Context, tx *sql.Tx, id string) (*PIC, error)
		DeletePIC(ctx context.Context, tx *sql.Tx, id string)
	}

	PICService interface {
		SavePIC(ctx context.Context, request *SaveRequestPIC) PICResponse
		UpdatePIC(ctx context.Context, request *UpdateRequestPIC) PICResponse
		GetPICs(ctx context.Context) []PICResponse
		GetPICsWithPagination(ctx context.Context, params *helper.PaginationParams) []PICResponse
		GetPIC(ctx context.Context, id string) PICResponse
		DeletePIC(ctx context.Context, id string)
	}

	PICHandler interface {
		SavePIC() http.HandlerFunc
		UpdatePIC() http.HandlerFunc
		GetPICs() http.HandlerFunc
		GetPICsWithPagination() http.HandlerFunc
		GetPIC() http.HandlerFunc
		DeletePIC() http.HandlerFunc
	}
)

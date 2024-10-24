package model

import (
	"context"
	"database/sql"
	"net/http"
)

type (
	MenuGroup struct {
		Id   string
		Name string
	}

	MenuGroupResult struct {
		Id     string
		Name   string
		Status string
	}

	MenuGroupResponse struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	}

	MenuGroupRequest struct {
		Id   string `json:"id" validate:"max=5"`
		Name string `json:"name" validate:"required,min=1,max=50"`
	}

	MenuGroupRepository interface {
		GenerateMenuGroupKode(ctx context.Context, tx *sql.Tx) *string
		CreateMenuGroup(ctx context.Context, tx *sql.Tx, data *MenuGroup) *MenuGroup
		UpdateMenuGroup(ctx context.Context, tx *sql.Tx, data *MenuGroup) *MenuGroup
		GetMenuGroup(ctx context.Context, tx *sql.Tx, id string) (*MenuGroup, error)
		GetAllMenuGroups(ctx context.Context, tx *sql.Tx) *[]MenuGroupResult
		DeleteMenuGroup(ctx context.Context, tx *sql.Tx, id string)
	}

	MenuGroupService interface {
		StoreMenuGroup(ctx context.Context, request *MenuGroupRequest) MenuGroupResponse
		UpdateMenuGroup(ctx context.Context, request *MenuGroupRequest) MenuGroupResponse
		GetAllMenuGroups(ctx context.Context) []MenuGroupResponse
		GetOneMenuGroup(ctx context.Context, id string) MenuGroupResponse
		DeleteMenuGroup(ctx context.Context, id string)
	}

	MenuGroupHandler interface {
		SaveMenuGroup() http.HandlerFunc
		UpdateMenuGroup() http.HandlerFunc
		GetAllMenuGroups() http.HandlerFunc
		GetOneMenuGroup() http.HandlerFunc
		DeleteMenuGroup() http.HandlerFunc
	}
)

package model

import (
	"context"
	"database/sql"
)

type (
	Role struct {
		ID   int
		Name string
	}

	RoleResponse struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	RoleRepository interface {
		GetRoleByID(ctx context.Context, tx *sql.Tx, id string) (*Role, error)
	}

	RoleService interface {
		GetRoleByID(ctx context.Context, id string) RoleResponse
	}

	RoleHandler interface{}
)

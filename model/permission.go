package model

import (
	"context"
	"database/sql"
)

type (
	Permission struct {
		ID   int
		Name string
	}

	PermissionRole struct {
		IDRole       int
		IDPermission int
	}

	PermissionRoleResponse struct {
		IDRole       int `json:"id_role"`
		IDPermission int `json:"id_permission"`
	}

	PermissionRepository interface {
		GetRolePermissionsByID(ctx context.Context, tx *sql.Tx, id int) *[]PermissionRole
	}

	PermissionService interface {
		GetRolePermissionsByID(ctx context.Context, id int) []PermissionRoleResponse
	}

	PermissionHandler interface{}
)

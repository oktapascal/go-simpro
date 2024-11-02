package permission

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) GetRolePermissionsByID(ctx context.Context, tx *sql.Tx, id int) *[]model.PermissionRole {
	query := "select role_id,permission_id from permissions_roles where role_id=?"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var rolePermissions []model.PermissionRole
	for rows.Next() {
		var rolePermission model.PermissionRole

		err = rows.Scan(&rolePermission.IDRole, &rolePermission.IDPermission)
		if err != nil {
			panic(err)
		}

		rolePermissions = append(rolePermissions, rolePermission)
	}

	return &rolePermissions
}

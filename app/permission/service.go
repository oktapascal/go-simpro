package permission

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
)

type Service struct {
	rpo model.PermissionRepository
	db  *sql.DB
}

func (svc *Service) GetRolePermissionsByID(ctx context.Context, id int) []model.PermissionRoleResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	rolePermissions := svc.rpo.GetRolePermissionsByID(ctx, tx, id)

	var results []model.PermissionRoleResponse
	for _, value := range *rolePermissions {
		rolePermission := model.PermissionRoleResponse{
			IDRole:       value.IDRole,
			IDPermission: value.IDPermission,
		}

		results = append(results, rolePermission)
	}

	return results
}

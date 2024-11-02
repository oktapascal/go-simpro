package role

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/exception"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
)

type Service struct {
	rpo model.RoleRepository
	db  *sql.DB
}

func (svc Service) GetRoleByID(ctx context.Context, id string) model.RoleResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	role, err := svc.rpo.GetRoleByID(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.RoleResponse{
		ID:   role.ID,
		Name: role.Name,
	}
}

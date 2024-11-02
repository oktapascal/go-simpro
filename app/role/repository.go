package role

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) GetRoleByID(ctx context.Context, tx *sql.Tx, id string) (*model.Role, error) {
	query := "select id,name from roles where id=?"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	role := new(model.Role)
	if rows.Next() {
		err = rows.Scan(&role.ID, &role.Name)
		if err != nil {
			panic(err)
		}

		return role, nil
	}

	return role, errors.New("role not found")
}

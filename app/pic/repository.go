package pic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) SavePIC(ctx context.Context, tx *sql.Tx, data *model.PIC) {
	query := "insert into pics (id,name,email,phone) values (?,?,?,?)"

	_, err := tx.ExecContext(ctx, query, data.ID, data.Name, data.Email, data.Phone)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) UpdatePIC(ctx context.Context, tx *sql.Tx, data *model.PIC) {
	query := "update pics set name=?,email=?,phone=?,updated_at=current_timestamp where id=?"

	_, err := tx.ExecContext(ctx, query, data.Name, data.Email, data.Phone, data.ID)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) GetPICs(ctx context.Context, tx *sql.Tx) *[]model.PIC {
	query := `select id, name, phone, email, 
    case
		when timestampdiff(minute, created_at, now()) <= 10 then 'CREATED'
		when timestampdiff(minute, updated_at, now()) <= 10 then 'UPDATED'
		else 'NONE' 
	end as status
	from pics 
	where deleted_at is null
	order by created_at asc, updated_at desc`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var pics []model.PIC
	for rows.Next() {
		var pic model.PIC
		err = rows.Scan(&pic.ID, &pic.Name, &pic.Phone, &pic.Email, &pic.Status)
		if err != nil {
			panic(err)
		}

		pics = append(pics, pic)
	}

	return &pics
}

func (rpo *Repository) GetPICsWithPagination(ctx context.Context, tx *sql.Tx, params *helper.PaginationParams) *[]model.PIC {
	query := "select t1.id, t1.name, t1.email, t1.phone from pics t1 where t1.deleted_at is null"
	var args []any

	// if there is a filter, add a filter condition
	if params.FilterBy != "" && params.FilterValue != "" {
		query += fmt.Sprintf(" and %s=?", params.FilterBy)
		args = append(args, params.FilterValue)
	}

	// add sorting by and order by
	if params.SortBy != "" {
		query += fmt.Sprintf(" order by %s %s", params.SortBy, params.OrderBy)
	} else {
		// if there is no sort by, give default sort by
		query += "order by id asc"
	}

	// add limit and offset for paginate
	offset := (params.Page - 1) * params.PageSize
	query += " limit ? offset ?"
	args = append(args, params.PageSize, offset)

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var pics []model.PIC
	for rows.Next() {
		var pic model.PIC
		err = rows.Scan(&pic.ID, &pic.Name, &pic.Email, &pic.Phone)
		if err != nil {
			panic(err)
		}

		pics = append(pics, pic)
	}

	return &pics
}

func (rpo *Repository) GetPIC(ctx context.Context, tx *sql.Tx, id string) (*model.PIC, error) {
	query := "select id,name,phone,email from pics where id=?"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	pic := new(model.PIC)
	if rows.Next() {
		err = rows.Scan(&pic.ID, &pic.Name, &pic.Phone, &pic.Email)
		if err != nil {
			panic(err)
		}

		return pic, nil
	}

	return nil, errors.New("pic not found")
}

func (rpo *Repository) DeletePIC(ctx context.Context, tx *sql.Tx, id string) {
	query := "update pics set deleted_at=current_timestamp where id=?"

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
}

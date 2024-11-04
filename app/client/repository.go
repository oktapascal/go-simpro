package client

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"strings"
)

type Repository struct{}

func (rpo *Repository) SaveClient(ctx context.Context, tx *sql.Tx, data *model.Client) {
	query := "insert into clients (id,name,address,phone) values (?,?,?,?)"

	_, err := tx.ExecContext(ctx, query, data.ID, data.Name, data.Address, data.Phone)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) SaveClientPIC(ctx context.Context, tx *sql.Tx, data *[]model.ClientPIC) {
	placeholder := ""

	var args []any

	for i, row := range *data {
		placeholder += "(?, ?, ?, ?, ?)"

		if i < len(*data)-1 {
			placeholder += ","
		}

		args = append(args, row.IDClient, row.Name, row.Phone, row.Email, row.Address)
	}

	query := fmt.Sprintf("insert into clients_pic (client_id,name,phone,email,address) values %s", placeholder)

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) UpdateClient(ctx context.Context, tx *sql.Tx, data *model.Client) {
	query := "update clients set name=?,address=?,phone=? where id=?"

	_, err := tx.ExecContext(ctx, query, data.Name, data.Address, data.Phone, data.ID)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) UpdateClientPIC(ctx context.Context, tx *sql.Tx, data *[]model.ClientPIC) {
	query := "update clients_pic set name=?, phone=?, email=?, address=? where id=? and client_id=?"

	stmt, err := tx.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for _, value := range *data {
		_, err := stmt.ExecContext(ctx, value.Name, value.Phone, value.Email, value.Address, value.ID, value.IDClient)
		if err != nil {
			panic(err)
		}
	}
}

func (rpo *Repository) GetClients(ctx context.Context, tx *sql.Tx) *[]model.Client {
	query := `select id, name, phone, address, 
    case
		when timestampdiff(minute, created_at, now()) <= 10 then 'CREATED'
		when timestampdiff(minute, updated_at, now()) <= 10 then 'UPDATED'
		else 'NONE' 
	end as status
	from clients 
	where deleted_at is null
	order by created_at asc, updated_at desc`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var clients []model.Client
	for rows.Next() {
		var client model.Client
		err = rows.Scan(&client.ID, &client.Name, &client.Phone, &client.Address, &client.Status)
		if err != nil {
			panic(err)
		}

		clients = append(clients, client)
	}

	return &clients
}

func (rpo *Repository) GetClientsWithPagination(ctx context.Context, tx *sql.Tx, params *helper.PaginationParams) *[]model.Client {
	query := "select t1.id, t1.name, t1.address, t1.phone from clients t1 where t1.deleted_at is null"
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

	var clients []model.Client
	for rows.Next() {
		var client model.Client
		err = rows.Scan(&client.ID, &client.Name, &client.Address, &client.Phone)
		if err != nil {
			panic(err)
		}

		clients = append(clients, client)
	}

	return &clients
}

func (rpo *Repository) GetClient(ctx context.Context, tx *sql.Tx, id string) (*model.Client, error) {
	query := "select id,name,phone,address from clients where id=?"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	client := new(model.Client)
	if rows.Next() {
		err = rows.Scan(&client.ID, &client.Name, &client.Phone, &client.Address)
		if err != nil {
			panic(err)
		}

		return client, nil
	}

	return nil, errors.New("client not found")
}

func (rpo *Repository) GetClientPIC(ctx context.Context, tx *sql.Tx, id string) *[]model.ClientPIC {
	query := "select id, name, phone, email, address from clients_pic where client_id = ? and deleted_at is null"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var clientsPic []model.ClientPIC
	for rows.Next() {
		var clientPic model.ClientPIC
		err = rows.Scan(&clientPic.ID, &clientPic.Name, &clientPic.Phone, &clientPic.Email, &clientPic.Address)
		if err != nil {
			panic(err)
		}

		clientsPic = append(clientsPic, clientPic)
	}

	return &clientsPic
}

func (rpo *Repository) DeleteClient(ctx context.Context, tx *sql.Tx, id string) {
	query := "update clients set deleted_at=current_timestamp where id=?"

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	query = "update clients_pic set deleted_at=current_timestamp where client_id=? and deleted_at is null"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) DeleteClientPIC(ctx context.Context, tx *sql.Tx, IDClient string, id []int) {
	placeholders := make([]string, len(id))
	for i := range id {
		placeholders[i] = "?"
	}

	query := fmt.Sprintf("update clients_pic set deleted_at=current_timestamp where client_id=? and id not in (%s)", strings.Join(placeholders, ","))

	args := make([]any, len(id)+1)
	args[0] = IDClient
	for i, value := range id {
		args[i+1] = value
	}

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		panic(err)
	}
}

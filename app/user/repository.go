package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oktapascal/go-simpro/model"
)

type Repository struct{}

func (rpo *Repository) GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (*model.User, error) {
	query := `select id,username,email,password,name,phone,role_id,menu_group_id,avatar,status_active
	from users where email=? and status_active=1`

	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	user := new(model.User)

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Name, &user.Phone, &user.IDRole,
			&user.IDMenuGroup, &user.Avatar, &user.StatusActive)
		if err != nil {
			panic(err)
		}

		return user, nil
	}

	return user, errors.New("user not found")
}

func (rpo *Repository) GetUserByUsername(ctx context.Context, tx *sql.Tx, username string) (*model.User, error) {
	query := `select id,username,email,password,name,phone,role_id,menu_group_id,avatar,status_active
	from users where username=? and status_active=1`

	rows, err := tx.QueryContext(ctx, query, username)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	user := new(model.User)

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Name, &user.Phone, &user.IDRole,
			&user.IDMenuGroup, &user.Avatar, &user.StatusActive)
		if err != nil {
			panic(err)
		}

		return user, nil
	}

	return user, errors.New("user not found")
}

func (rpo *Repository) GetUserByID(ctx context.Context, tx *sql.Tx, id string) (*model.User, error) {
	query := `select id,username,email,password,name,phone,role_id,menu_group_id,avatar,status_active
	from users where id=? and status_active=1`

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	user := new(model.User)

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Name, &user.Phone, &user.IDRole,
			&user.IDMenuGroup, &user.Avatar, &user.StatusActive)
		if err != nil {
			panic(err)
		}

		return user, nil
	}

	return user, errors.New("user not found")
}

func (rpo *Repository) SaveUser(ctx context.Context, tx *sql.Tx, data *model.User) {
	query := `insert into users (id,username,email,password,name,phone,role_id,menu_group_id,status_active,avatar) 
	values (?,?,?,?,?,?,?,?,?,?)`

	_, err := tx.ExecContext(ctx, query, data.ID, data.Username, data.Email, data.Password, data.Name, data.Phone,
		data.IDRole, data.IDMenuGroup, 1, "")
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) UpdateProfilePhotoUser(ctx context.Context, tx *sql.Tx, data *model.User) {
	query := "update users set avatar=? where id=?"

	_, err := tx.ExecContext(ctx, query, data.Avatar, data.ID)
	if err != nil {
		panic(err)
	}
}

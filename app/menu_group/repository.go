package menu_group

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/oktapascal/go-simpro/model"
	"strconv"
)

type Repository struct {
}

func (rpo *Repository) GenerateMenuGroupKode(ctx context.Context, tx *sql.Tx) *string {
	query := "select id from menu_groups order by created_at desc limit 1"

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	var id string
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}

		strNumber := id[3:]
		number, errConvert := strconv.Atoi(strNumber)
		if errConvert != nil {
			panic(errConvert)
		}

		number++
		strNumber = strconv.Itoa(number)

		if len(strNumber) == 2 {
			id = fmt.Sprintf("MG-%s", strNumber)
		} else {
			id = fmt.Sprintf("MG-0%s", strNumber)
		}
	} else {
		id = "MG-01"
	}

	return &id
}

func (rpo *Repository) CreateMenuGroup(ctx context.Context, tx *sql.Tx, data *model.MenuGroup) *model.MenuGroup {
	query := "insert into menu_groups (id, name) values (?,?)"

	_, err := tx.ExecContext(ctx, query, data.Id, data.Name)
	if err != nil {
		panic(err)
	}

	return data
}

func (rpo *Repository) UpdateMenuGroup(ctx context.Context, tx *sql.Tx, data *model.MenuGroup) *model.MenuGroup {
	query := "update menu_groups set name=? where id=?"

	_, err := tx.ExecContext(ctx, query, data.Name, data.Id)
	if err != nil {
		panic(err)
	}

	return data
}

func (rpo *Repository) GetMenuGroup(ctx context.Context, tx *sql.Tx, id string) (*model.MenuGroup, error) {
	query := "select id, name from menu_groups where id=?"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	menuGroup := new(model.MenuGroup)
	if rows.Next() {
		err = rows.Scan(&menuGroup.Id, &menuGroup.Name)
		if err != nil {
			panic(err)
		}

		return menuGroup, nil
	} else {
		return nil, errors.New("menu group not found")
	}
}

func (rpo *Repository) GetAllMenuGroups(ctx context.Context, tx *sql.Tx) *[]model.MenuGroupResult {
	query := `select id, name,
    case
		when timestampdiff(minute, created_at, now()) <= 10 then 'CREATED'
		when timestampdiff(minute, updated_at, now()) <= 10 then 'UPDATED'
		else 'NONE' 
	end as status
	from menu_groups 
	where deleted_at is null`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	var menuGroups []model.MenuGroupResult
	for rows.Next() {
		var menuGroup model.MenuGroupResult
		err = rows.Scan(&menuGroup.Id, &menuGroup.Name, &menuGroup.Status)
		if err != nil {
			panic(err)
		}

		menuGroups = append(menuGroups, menuGroup)
	}

	return &menuGroups
}

func (rpo *Repository) DeleteMenuGroup(ctx context.Context, tx *sql.Tx, id string) {
	query := "update menu_groups set deleted_at=current_timestamp where id=?"

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
}

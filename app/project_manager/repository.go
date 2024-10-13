package project_manager

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

func (rpo *Repository) GenerateProjectManagerKode(ctx context.Context, tx *sql.Tx) *string {
	query := "select id from project_managers order by created_at desc limit 1"

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

		strNumber := id[4:]
		number, errConvert := strconv.Atoi(strNumber)
		if errConvert != nil {
			panic(errConvert)
		}

		number++
		strNumber = strconv.Itoa(number)

		if len(strNumber) == 3 {
			id = fmt.Sprintf("PIC-%s", strNumber)
		} else if len(strNumber) == 2 {
			id = fmt.Sprintf("PIC-0%s", strNumber)
		} else {
			id = fmt.Sprintf("PIC-00%s", strNumber)
		}
	} else {
		id = "PIC-001"
	}

	return &id
}

func (rpo *Repository) CreateProjectManager(ctx context.Context, tx *sql.Tx, data *model.ProjectManager) *model.ProjectManager {
	query := "insert into project_managers (id, name) values (?, ?)"

	_, err := tx.ExecContext(ctx, query, data.Id, data.Name)
	if err != nil {
		panic(err)
	}

	return data
}

func (rpo *Repository) UpdateProjectManager(ctx context.Context, tx *sql.Tx, data *model.ProjectManager) *model.ProjectManager {
	query := "update project_managers set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, query, data.Name, data.Id)
	if err != nil {
		panic(err)
	}

	return data
}

func (rpo *Repository) GetProjectManagersNoPagination(ctx context.Context, tx *sql.Tx) *[]model.ProjectManagerResult {
	query := `select id, name, 
    case
		when timestampdiff(minute, created_at, now()) <= 10 then 'CREATED'
		when timestampdiff(minute, updated_at, now()) <= 10 then 'UPDATED'
		else 'NONE' 
	end as status
	from project_managers 
	where deleted_at is null
	order by created_at asc, updated_at desc`

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

	var projectManagers []model.ProjectManagerResult
	for rows.Next() {
		var projectManager model.ProjectManagerResult
		err = rows.Scan(&projectManager.Id, &projectManager.Name, &projectManager.Status)
		if err != nil {
			panic(err)
		}

		projectManagers = append(projectManagers, projectManager)
	}

	return &projectManagers
}

func (rpo *Repository) GetProjectManager(ctx context.Context, tx *sql.Tx, id string) (*model.ProjectManager, error) {
	query := "select id, name from project_managers where id = ?"

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

	projectManager := new(model.ProjectManager)
	if rows.Next() {
		err = rows.Scan(&projectManager.Id, &projectManager.Name)
		if err != nil {
			panic(err)
		}

		return projectManager, nil
	} else {
		return nil, errors.New("project manager not found")
	}
}

func (rpo *Repository) DeleteProjectManager(ctx context.Context, tx *sql.Tx, id string) {
	query := "update project_managers set deleted_at = current_timestamp where id = ?"

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
}

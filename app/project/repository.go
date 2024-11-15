package project

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/oktapascal/go-simpro/model"
	"strings"
)

type Repository struct{}

func (rpo *Repository) SaveProject(ctx context.Context, tx *sql.Tx, data *model.Project) {
	query := "insert into projects (id,client_id,client_pic_id,description,project_type,project_status) values (?,?,?,?,?,?)"

	_, err := tx.ExecContext(ctx, query, data.ID, data.IDClient, data.IDClientPIC, data.Description, data.ProjectType, data.ProjectStatus)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) UpdateProject(ctx context.Context, tx *sql.Tx, data *model.Project) {
	query := "update projects set client_id=?,client_pic_id=?,description=?,project_type=?,updated_at=current_timestamp where id=?"

	_, err := tx.ExecContext(ctx, query, data.IDClient, data.IDClientPIC, data.Description, data.ProjectType, data.ID)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) GetProjects(ctx context.Context, tx *sql.Tx) *[]model.Project {
	query := `select id, description, project_type, project_status, 
    case
		when timestampdiff(minute, created_at, now()) <= 10 then 'CREATED'
		when timestampdiff(minute, updated_at, now()) <= 10 then 'UPDATED'
		else 'NONE' 
	end as status
	from projects 
	where closed_at is null and project_status = 'OPEN'
	order by created_at asc, updated_at desc`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var project model.Project
		err = rows.Scan(&project.ID, &project.Description, &project.ProjectType, &project.ProjectStatus, &project.Status)
		if err != nil {
			panic(err)
		}

		projects = append(projects, project)
	}

	return &projects
}

func (rpo *Repository) GetProject(ctx context.Context, tx *sql.Tx, id string) (*model.Project, error) {
	query := `select t1.id,t1.client_id,t1.client_pic_id,t1.description,t1.project_type,t1.project_status,t2.name client_name,t3.name client_pic_name
	from projects t1
	inner join clients t2 on t1.client_id=t2.id
	inner join clients_pic t3 on t1.client_pic_id=t3.id
	where t1.id=?`

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	project := new(model.Project)
	if rows.Next() {
		err = rows.Scan(&project.ID, &project.IDClient, &project.IDClientPIC, &project.Description, &project.ProjectType,
			&project.ProjectStatus, &project.Client.Name, &project.ClientPIC.Name)
		if err != nil {
			panic(err)
		}

		return project, nil
	}

	return nil, errors.New("project not found")
}

func (rpo *Repository) SaveCloseProject(ctx context.Context, tx *sql.Tx, data *model.Project) {
	query := "update projects set description_close=?,project_status=?,closed_at=current_timestamp where id=?"

	_, err := tx.ExecContext(ctx, query, data.DescriptionClosed, data.ProjectStatus, data.ID)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) SaveCloseProjectDoc(ctx context.Context, tx *sql.Tx, data *[]model.ProjectDoc) {
	placeholder := ""

	var args []any

	for i, row := range *data {
		placeholder += "(?, ?, ?)"

		if i < len(*data)-1 {
			placeholder += ","
		}

		args = append(args, row.IDProject, row.Description, row.FileName)
	}

	query := fmt.Sprintf("insert into project_docs (project_id,description,file_name) values %s", placeholder)

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) UpdateCloseProject(ctx context.Context, tx *sql.Tx, data *model.Project) {
	query := "update projects set description_close=?,project_status=?,updated_at=current_timestamp where id=?"

	_, err := tx.ExecContext(ctx, query, data.DescriptionClosed, data.ProjectStatus, data.ID)
	if err != nil {
		panic(err)
	}
}

func (rpo *Repository) UpdateCloseProjectDoc(ctx context.Context, tx *sql.Tx, data *[]model.ProjectDoc) {
	query := "update project_docs set description=?, file_name=?, updated_at=current_timestamp where id=? and project_id=?"

	stmt, err := tx.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for _, value := range *data {
		_, err := stmt.ExecContext(ctx, value.Description, value.FileName, value.ID, value.IDProject)
		if err != nil {
			panic(err)
		}
	}
}

func (rpo *Repository) GetCloseProjects(ctx context.Context, tx *sql.Tx) *[]model.Project {
	query := `select id, description_close, project_type, project_status, 
    case
		when timestampdiff(minute, created_at, now()) <= 10 then 'CREATED'
		when timestampdiff(minute, updated_at, now()) <= 10 then 'UPDATED'
		else 'NONE' 
	end as status
	from projects 
	where closed_at is not null and project_status != 'OPEN'
	order by created_at asc, updated_at desc`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var project model.Project
		err = rows.Scan(&project.ID, &project.DescriptionClosed, &project.ProjectType, &project.ProjectStatus, &project.Status)
		if err != nil {
			panic(err)
		}

		projects = append(projects, project)
	}

	return &projects
}

func (rpo *Repository) GetCloseProject(ctx context.Context, tx *sql.Tx, id string) (*model.Project, error) {
	query := `select t1.id,t1.client_id,t1.client_pic_id,t1.description,t1.project_type,t1.project_status,t1.description_close,
    t2.name client_name,t3.name client_pic_name
	from projects t1
	inner join clients t2 on t1.client_id=t2.id
	inner join clients_pic t3 on t1.client_pic_id=t3.id
	where t1.id=? and t1.project_status != 'OPEN' and t1.closed_at is not null`

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	project := new(model.Project)
	if rows.Next() {
		err = rows.Scan(&project.ID, &project.IDClient, &project.IDClientPIC, &project.Description, &project.ProjectType,
			&project.ProjectStatus, &project.DescriptionClosed, &project.Client.Name, &project.ClientPIC.Name)
		if err != nil {
			panic(err)
		}

		return project, nil
	}

	return nil, errors.New("project not found")
}

func (rpo *Repository) GetCloseProjectDoc(ctx context.Context, tx *sql.Tx, id string) *[]model.ProjectDoc {
	query := "select id,description,file_name from project_docs where project_id = ? and deleted_at is null"

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var projectDocs []model.ProjectDoc
	for rows.Next() {
		var projectDoc model.ProjectDoc
		err = rows.Scan(&projectDoc.ID, projectDoc.Description, &projectDoc.FileName)
		if err != nil {
			panic(err)
		}

		projectDocs = append(projectDocs, projectDoc)
	}

	return &projectDocs
}

func (rpo *Repository) DeleteCloseProjectDoc(ctx context.Context, tx *sql.Tx, IDProject string, id []int) {
	placeholders := make([]string, len(id))
	for i := range id {
		placeholders[i] = "?"
	}

	query := fmt.Sprintf("update project_docs set deleted_at=current_timestamp where project_id=? and id not in (%s)", strings.Join(placeholders, ","))

	args := make([]any, len(id)+1)
	args[0] = IDProject
	for i, value := range id {
		args[i+1] = value
	}

	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		panic(err)
	}
}

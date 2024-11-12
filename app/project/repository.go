package project

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oktapascal/go-simpro/model"
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
	query := "update projects set client_id=?,client_pic_id=?,description=?,project_type=? where id=?"

	_, err := tx.ExecContext(ctx, query, data.IDClient, data.IDClientPIC, data.Description, data.ID)
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
	where closed_at is null
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
	query := `select t1.id,t1.client_id,t1.client_pic_id,t1.description,t1.project_type,t2.name client_name,t3.name client_pic_name
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
			&project.Client.Name, &project.ClientPIC.Name)
		if err != nil {
			panic(err)
		}

		return project, nil
	}

	return nil, errors.New("project not found")
}

func (rpo *Repository) SaveCloseProject(ctx context.Context, tx *sql.Tx, data *model.Project) {
	//TODO implement me
	panic("implement me")
}

func (rpo *Repository) SaveCloseProjectDoc(ctx context.Context, tx *sql.Tx, data *[]model.ProjectDoc) {
	//TODO implement me
	panic("implement me")
}

func (rpo *Repository) UpdateCloseProject(ctx context.Context, tx *sql.Tx, data *model.Project) {
	//TODO implement me
	panic("implement me")
}

func (rpo *Repository) UpdateCloseProjectDoc(ctx context.Context, tx *sql.Tx, data *[]model.ProjectDoc) {
	//TODO implement me
	panic("implement me")
}

func (rpo *Repository) GetCloseProjects(ctx context.Context, tx *sql.Tx) *[]model.Project {
	//TODO implement me
	panic("implement me")
}

func (rpo *Repository) GetCloseProject(ctx context.Context, tx *sql.Tx, id string) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (rpo *Repository) DeleteCloseProjectDoc(ctx context.Context, tx *sql.Tx, IDProject string, id []int) {
	//TODO implement me
	panic("implement me")
}

package project

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/oktapascal/go-simpro/exception"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"time"
)

type Service struct {
	rpo model.ProjectRepository
	db  *sql.DB
}

func (svc *Service) SaveProject(ctx context.Context, request *model.SaveRequestProject) model.ProjectResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	var projectID string

	getProjectID := func() string {
		now := time.Now()
		id := now.Unix()

		return fmt.Sprintf("PRJ-%d", id)
	}

	projectID = getProjectID()

	_, err = svc.rpo.GetProject(ctx, tx, projectID)
	if err == nil {
		projectID = getProjectID()
	}

	project := new(model.Project)
	project.ID = projectID
	project.Description = request.Description
	project.ProjectStatus = "OPEN"
	project.ProjectType = request.ProjectType
	project.IDClient = request.IDClient
	project.IDClientPIC = request.IDClientPIC

	svc.rpo.SaveProject(ctx, tx, project)

	return model.ProjectResponse{
		ID:            projectID,
		Description:   project.Description,
		ProjectType:   project.ProjectType,
		ProjectStatus: project.ProjectStatus,
	}
}

func (svc *Service) UpdateProject(ctx context.Context, request *model.UpdateRequestProject) model.ProjectResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	project, err := svc.rpo.GetProject(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	project.Description = request.Description
	project.ProjectType = request.ProjectType
	project.IDClient = request.IDClient
	project.IDClientPIC = request.IDClientPIC

	svc.rpo.SaveProject(ctx, tx, project)

	return model.ProjectResponse{
		ID:            project.ID,
		Description:   project.Description,
		ProjectType:   project.ProjectType,
		ProjectStatus: project.ProjectStatus,
	}
}

func (svc *Service) GetProjects(ctx context.Context) []model.ProjectResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	projects := svc.rpo.GetProjects(ctx, tx)

	var result []model.ProjectResponse
	for _, value := range *projects {
		project := model.ProjectResponse{
			ID:            value.ID,
			Description:   value.Description,
			ProjectType:   value.ProjectType,
			ProjectStatus: value.ProjectStatus,
			Status:        value.Status,
		}

		result = append(result, project)
	}

	return result
}

func (svc *Service) GetProject(ctx context.Context, id string) model.ProjectResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	project, err := svc.rpo.GetProject(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ProjectResponse{
		ID:            project.ID,
		Description:   project.Description,
		ProjectType:   project.ProjectType,
		ProjectStatus: project.ProjectStatus,
		Status:        project.Status,
	}
}

func (svc *Service) SaveCloseProject(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (svc *Service) UpdateCloseProject(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (svc *Service) GetCloseProjects(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (svc *Service) GetCloseProject(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

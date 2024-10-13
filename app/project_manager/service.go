package project_manager

import (
	"context"
	"database/sql"
	"github.com/oktapascal/go-simpro/exception"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
)

type Service struct {
	rpo model.ProjectManagerRepository
	db  *sql.DB
}

func (svc *Service) StoreProjectManager(ctx context.Context, request *model.SaveProjectManagerRequest) model.ProjectManagerResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	id := svc.rpo.GenerateProjectManagerKode(ctx, tx)

	projectManager := new(model.ProjectManager)
	projectManager.Id = *id
	projectManager.Name = request.Name

	projectManager = svc.rpo.CreateProjectManager(ctx, tx, projectManager)

	return model.ProjectManagerResponse{
		Id:   projectManager.Id,
		Name: projectManager.Name,
	}
}

func (svc *Service) UpdateProjectManager(ctx context.Context, request *model.UpdateProjectManagerRequest) model.ProjectManagerResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	projectManager, errClient := svc.rpo.GetProjectManager(ctx, tx, request.Id)
	if errClient != nil {
		panic(exception.NewNotFoundError(errClient.Error()))
	}

	projectManager.Name = request.Name

	return model.ProjectManagerResponse{
		Id:   projectManager.Id,
		Name: projectManager.Name,
	}
}

func (svc *Service) GetProjectManagersNoPagination(ctx context.Context) []model.ProjectManagerResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	projectManagers := svc.rpo.GetProjectManagersNoPagination(ctx, tx)

	var result []model.ProjectManagerResponse
	if len(*projectManagers) > 0 {
		for _, value := range *projectManagers {
			projectManager := model.ProjectManagerResponse{
				Id:     value.Id,
				Name:   value.Name,
				Status: value.Status,
			}

			result = append(result, projectManager)
		}
	}

	return result
}

func (svc *Service) GetOneProjectManager(ctx context.Context, id string) model.ProjectManagerResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	projectManager, errProjectManager := svc.rpo.GetProjectManager(ctx, tx, id)
	if errProjectManager != nil {
		panic(exception.NewNotFoundError(errProjectManager.Error()))
	}

	return model.ProjectManagerResponse{
		Id:   projectManager.Id,
		Name: projectManager.Name,
	}
}

func (svc *Service) DeleteProjectManager(ctx context.Context, id string) {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	_, errProjectManager := svc.rpo.GetProjectManager(ctx, tx, id)
	if errProjectManager != nil {
		panic(exception.NewNotFoundError(errProjectManager.Error()))
	}

	svc.rpo.DeleteProjectManager(ctx, tx, id)
}

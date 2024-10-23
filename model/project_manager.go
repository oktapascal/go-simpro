package model

import (
	"context"
	"database/sql"
	"net/http"
)

type (
	ProjectManager struct {
		Id    string
		Name  string
		Email string
		Phone string
	}

	ProjectManagerResult struct {
		Id     string
		Name   string
		Status string
	}

	SaveProjectManagerRequest struct {
		Name  string `json:"name" validate:"required,min=1,max=50"`
		Email string `json:"email" validate:"required,email,min=1,max=50"`
		Phone string `json:"phone" validate:"required,min=1,max=13"`
	}

	UpdateProjectManagerRequest struct {
		Id    string `json:"id"`
		Name  string `json:"name" validate:"required,min=1,max=50"`
		Email string `json:"email" validate:"required,email,min=1,max=50"`
		Phone string `json:"phone" validate:"required,min=1,max=13"`
	}

	ProjectManagerResponse struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Phone  string `json:"phone"`
		Status string `json:"status"`
	}

	ProjectManagerRepository interface {
		GenerateProjectManagerKode(ctx context.Context, tx *sql.Tx) *string
		CreateProjectManager(ctx context.Context, tx *sql.Tx, data *ProjectManager) *ProjectManager
		UpdateProjectManager(ctx context.Context, tx *sql.Tx, data *ProjectManager) *ProjectManager
		GetProjectManagersNoPagination(ctx context.Context, tx *sql.Tx) *[]ProjectManagerResult
		GetProjectManager(ctx context.Context, tx *sql.Tx, id string) (*ProjectManager, error)
		DeleteProjectManager(ctx context.Context, tx *sql.Tx, id string)
	}

	ProjectManagerService interface {
		StoreProjectManager(ctx context.Context, request *SaveProjectManagerRequest) ProjectManagerResponse
		UpdateProjectManager(ctx context.Context, request *UpdateProjectManagerRequest) ProjectManagerResponse
		GetProjectManagersNoPagination(ctx context.Context) []ProjectManagerResponse
		GetOneProjectManager(ctx context.Context, id string) ProjectManagerResponse
		DeleteProjectManager(ctx context.Context, id string)
	}

	ProjectManagerHandler interface {
		SaveProjectManager() http.HandlerFunc
		UpdateProjectManager() http.HandlerFunc
		GetProjectManagersNoPagination() http.HandlerFunc
		GetOneProjectManager() http.HandlerFunc
		DeleteProjectManager() http.HandlerFunc
	}
)

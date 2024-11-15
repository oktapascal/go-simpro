package model

import (
	"context"
	"database/sql"
	"net/http"
)

type (
	Project struct {
		ID            string
		IDClient      string
		IDClientPIC   int
		Description   string
		ProjectType   string
		ProjectStatus string
		Status        string
		Client        Client
		ClientPIC     ClientPIC
	}

	ProjectDoc struct {
		ID          int
		IDProject   string
		Description string
		FileName    string
	}

	SaveRequestProjectDoc struct {
		Description string `json:"description" validate:"required,min=1,max=250"`
		FileName    string `json:"file_name" validate:"required,min=1,max=20"`
	}

	SaveRequestProject struct {
		IDClient    string `json:"id_client" validate:"required,min=1,max=14"`
		IDClientPIC int    `json:"id_client_pic" validate:"required"`
		Description string `json:"description" validate:"required,min=1,max=250"`
		ProjectType string `json:"project_type" validate:"required,min=1,max=20"`
	}

	UpdateRequestProject struct {
		ID          string `json:"id" validate:"required"`
		IDClient    string `json:"id_client" validate:"required,min=1,max=14"`
		IDClientPIC int    `json:"id_client_pic" validate:"required"`
		Description string `json:"description" validate:"required,min=1,max=250"`
		ProjectType string `json:"project_type" validate:"required,min=1,max=20"`
	}

	ProjectResponse struct {
		ID            string `json:"id"`
		Description   string `json:"description"`
		ProjectType   string `json:"project_type"`
		ProjectStatus string `json:"project_status"`
		Status        string `json:"status"`
	}

	ProjectRepository interface {
		SaveProject(ctx context.Context, tx *sql.Tx, data *Project)
		UpdateProject(ctx context.Context, tx *sql.Tx, data *Project)
		GetProjects(ctx context.Context, tx *sql.Tx) *[]Project
		GetProject(ctx context.Context, tx *sql.Tx, id string) (*Project, error)
		SaveCloseProject(ctx context.Context, tx *sql.Tx, data *Project)
		SaveCloseProjectDoc(ctx context.Context, tx *sql.Tx, data *[]ProjectDoc)
		UpdateCloseProject(ctx context.Context, tx *sql.Tx, data *Project)
		UpdateCloseProjectDoc(ctx context.Context, tx *sql.Tx, data *[]ProjectDoc)
		GetCloseProjects(ctx context.Context, tx *sql.Tx) *[]Project
		GetCloseProject(ctx context.Context, tx *sql.Tx, id string) (*Project, error)
		DeleteCloseProjectDoc(ctx context.Context, tx *sql.Tx, IDProject string, id []int)
	}

	ProjectService interface {
		SaveProject(ctx context.Context, request *SaveRequestProject) ProjectResponse
		UpdateProject(ctx context.Context, request *UpdateRequestProject) ProjectResponse
		GetProjects(ctx context.Context) []ProjectResponse
		GetProject(ctx context.Context, id string) ProjectResponse
		SaveCloseProject(ctx context.Context)
		UpdateCloseProject(ctx context.Context)
		GetCloseProjects(ctx context.Context)
		GetCloseProject(ctx context.Context)
	}

	ProjectHandler interface {
		SaveProject() http.HandlerFunc
		UpdateProject() http.HandlerFunc
		GetProjects() http.HandlerFunc
		GetProject() http.HandlerFunc
		SaveCloseProject() http.HandlerFunc
		UpdateCloseProject() http.HandlerFunc
		GetCloseProjects() http.HandlerFunc
		GetCloseProject() http.HandlerFunc
	}
)

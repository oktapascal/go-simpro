package model

import (
	"context"
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type (
	User struct {
		ID           string
		Username     string
		Email        string
		Password     string
		Name         string
		Phone        string
		IDRole       int
		IDMenuGroup  string
		StatusActive bool
		Avatar       string
		UpdatedAt    string
		DeletedAt    string
	}

	UserResponse struct {
		ID           string `json:"id"`
		Username     string `json:"username"`
		Password     string `json:"-"`
		Email        string `json:"email"`
		Name         string `json:"name"`
		Phone        string `json:"phone"`
		IDRole       int    `json:"id_role"`
		IDMenuGroup  string `json:"id_menu_group"`
		StatusActive bool   `json:"status_active"`
		Avatar       string `json:"avatar"`
	}

	SaveRequestUser struct {
		Username             string `json:"username" validate:"required,min=1,max=50"`
		Email                string `json:"email" validate:"required,email,min=1,max=50"`
		Password             string `json:"password" validate:"required,min=1,max=50"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
		Name                 string `json:"name" validate:"required,min=1,max=50"`
		Phone                string `json:"phone" validate:"required,min=1,max=13"`
		IDRole               int    `json:"id_role" validate:"required"`
		IDMenuGroup          string `json:"id_menu_group"`
	}

	UpdateRequestUser struct {
		Email string `json:"email" validate:"required,email,min=1,max=50"`
		Name  string `json:"name" validate:"required,min=1,max=50"`
		Phone string `json:"phone" validate:"required,min=1,max=13"`
	}

	UserRepository interface {
		GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (*User, error)
		GetUserByUsername(ctx context.Context, tx *sql.Tx, username string) (*User, error)
		GetUserByID(ctx context.Context, tx *sql.Tx, id string) (*User, error)
		SaveUser(ctx context.Context, tx *sql.Tx, data *User)
		UpdateProfilePhotoUser(ctx context.Context, tx *sql.Tx, data *User)
		UpdateUser(ctx context.Context, tx *sql.Tx, data *User)
	}

	UserService interface {
		GetUserByEmail(ctx context.Context, email string) UserResponse
		GetUserByUsername(ctx context.Context, username string) UserResponse
		GetUserByID(ctx context.Context, id string) UserResponse
		SaveUser(ctx context.Context, request *SaveRequestUser) UserResponse
		UpdateProfilePhotoUser(ctx context.Context, fileName string, claims jwt.MapClaims) UserResponse
		UpdateUser(ctx context.Context, request *UpdateRequestUser, claims jwt.MapClaims) UserResponse
	}

	UserHandler interface {
		SaveUser() http.HandlerFunc
		UpdateProfilePhotoUser() http.HandlerFunc
		UpdateUser() http.HandlerFunc
	}
)

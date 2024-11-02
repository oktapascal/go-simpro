package user

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
	rpo model.UserRepository
	db  *sql.DB
}

func (svc *Service) GetUserByEmail(ctx context.Context, email string) model.UserResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	user, err := svc.rpo.GetUserByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Name:         user.Name,
		Phone:        user.Phone,
		IDRole:       user.IDRole,
		IDMenuGroup:  user.IDMenuGroup,
		StatusActive: user.StatusActive,
		Avatar:       user.Avatar,
	}
}

func (svc *Service) GetUserByUsername(ctx context.Context, username string) model.UserResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	user, err := svc.rpo.GetUserByUsername(ctx, tx, username)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Name:         user.Name,
		Phone:        user.Phone,
		IDRole:       user.IDRole,
		IDMenuGroup:  user.IDMenuGroup,
		StatusActive: user.StatusActive,
		Avatar:       user.Avatar,
	}
}

func (svc *Service) GetUserByID(ctx context.Context, id string) model.UserResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	user, err := svc.rpo.GetUserByID(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Name:         user.Name,
		Phone:        user.Phone,
		IDRole:       user.IDRole,
		IDMenuGroup:  user.IDMenuGroup,
		StatusActive: user.StatusActive,
		Avatar:       user.Avatar,
	}
}

func (svc *Service) SaveUser(ctx context.Context, request *model.SaveRequestUser) model.UserResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	var userID string

	getUserID := func() string {
		now := time.Now()
		id := now.Unix()

		return fmt.Sprintf("USR-%d", id)
	}

	userID = getUserID()

	_, err = svc.rpo.GetUserByID(ctx, tx, userID)
	if err == nil {
		userID = getUserID()
	}

	user := new(model.User)

	hash, errHash := helper.Hash(request.Password)
	if errHash == nil {
		user.Password = hash
	}

	user.ID = userID
	user.Username = request.Username
	user.Email = request.Email
	user.Name = request.Name
	user.Phone = request.Phone
	user.IDRole = request.IDRole
	user.IDMenuGroup = request.IDMenuGroup

	svc.rpo.SaveUser(ctx, tx, user)

	return model.UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Name:         user.Name,
		Phone:        user.Phone,
		IDRole:       user.IDRole,
		IDMenuGroup:  user.IDMenuGroup,
		StatusActive: true,
		Avatar:       user.Avatar,
	}
}

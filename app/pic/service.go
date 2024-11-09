package pic

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
	rpo model.PICRepository
	db  *sql.DB
}

func (svc *Service) SavePIC(ctx context.Context, request *model.SaveRequestPIC) model.PICResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	var PICID string

	getPICID := func() string {
		now := time.Now()
		id := now.Unix()

		return fmt.Sprintf("PIC-%d", id)
	}

	PICID = getPICID()

	_, err = svc.rpo.GetPIC(ctx, tx, PICID)
	if err == nil {
		PICID = getPICID()
	}

	pic := new(model.PIC)
	pic.ID = PICID
	pic.Name = request.Name
	pic.Email = request.Email
	pic.Phone = request.Phone

	svc.rpo.SavePIC(ctx, tx, pic)

	return model.PICResponse{
		ID:    PICID,
		Name:  pic.Name,
		Phone: pic.Phone,
		Email: pic.Email,
	}
}

func (svc *Service) UpdatePIC(ctx context.Context, request *model.UpdateRequestPIC) model.PICResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	pic, err := svc.rpo.GetPIC(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	pic.Name = request.Name
	pic.Email = request.Email
	pic.Phone = request.Phone

	svc.rpo.UpdatePIC(ctx, tx, pic)

	return model.PICResponse{
		ID:    pic.ID,
		Name:  pic.Name,
		Phone: pic.Phone,
		Email: pic.Email,
	}
}

func (svc *Service) GetPICs(ctx context.Context) []model.PICResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	pics := svc.rpo.GetPICs(ctx, tx)

	var result []model.PICResponse
	for _, value := range *pics {
		pic := model.PICResponse{
			ID:     value.ID,
			Name:   value.Name,
			Email:  value.Email,
			Phone:  value.Phone,
			Status: value.Status,
		}

		result = append(result, pic)
	}

	return result
}

func (svc *Service) GetPICsWithPagination(ctx context.Context, params *helper.PaginationParams) []model.PICResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	pics := svc.rpo.GetPICsWithPagination(ctx, tx, params)

	var result []model.PICResponse
	if len(*pics) > 0 {
		for _, value := range *pics {
			pic := model.PICResponse{
				ID:    value.ID,
				Name:  value.Name,
				Email: value.Email,
				Phone: value.Phone,
			}

			result = append(result, pic)
		}
	}

	return result
}

func (svc *Service) GetPIC(ctx context.Context, id string) model.PICResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	pic, err := svc.rpo.GetPIC(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.PICResponse{
		ID:    pic.ID,
		Name:  pic.Name,
		Phone: pic.Phone,
		Email: pic.Email,
	}
}

func (svc *Service) DeletePIC(ctx context.Context, id string) {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	_, err = svc.rpo.GetPIC(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	svc.rpo.DeletePIC(ctx, tx, id)
}

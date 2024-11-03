package client

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
	rpo model.ClientRepository
	db  *sql.DB
}

func (svc *Service) SaveClient(ctx context.Context, request *model.SaveRequestClient) model.ClientResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	var clientID string

	getClientID := func() string {
		now := time.Now()
		id := now.Unix()

		return fmt.Sprintf("CLN-%d", id)
	}

	clientID = getClientID()

	_, err = svc.rpo.GetClient(ctx, tx, clientID)
	if err == nil {
		clientID = getClientID()
	}

	client := new(model.Client)
	client.ID = clientID
	client.Name = request.Name
	client.Address = request.Address
	client.Phone = request.Phone

	svc.rpo.SaveClient(ctx, tx, client)

	var clientPIC []model.ClientPIC

	for _, value := range request.ClientPIC {
		clientsPIC := model.ClientPIC{
			IDClient: clientID,
			Name:     value.Name,
			Phone:    value.Phone,
			Email:    value.Email,
			Address:  value.Address,
		}

		clientPIC = append(clientPIC, clientsPIC)
	}

	svc.rpo.SaveClientPIC(ctx, tx, &clientPIC)

	return model.ClientResponse{
		ID:      clientID,
		Name:    client.Name,
		Address: client.Address,
		Phone:   client.Phone,
	}
}

func (svc *Service) UpdateClient(ctx context.Context, request *model.UpdateRequestClient) model.ClientResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	client, errClient := svc.rpo.GetClient(ctx, tx, request.ID)
	if errClient != nil {
		panic(exception.NewNotFoundError(errClient.Error()))
	}

	client.Name = request.Name
	client.Address = request.Address
	client.Phone = request.Phone

	var clientPicCollections []int

	for _, value := range request.ClientPIC {
		if value.ID != 0 {
			clientPicCollections = append(clientPicCollections, value.ID)
		}
	}

	svc.rpo.DeleteClientPIC(ctx, tx, request.ID, clientPicCollections)

	var clientsPICUpdate []model.ClientPIC
	var clientsPICInsert []model.ClientPIC

	for _, value := range request.ClientPIC {
		if value.ID != 0 {
			clientPICUpdate := model.ClientPIC{
				ID:       value.ID,
				IDClient: request.ID,
				Name:     value.Name,
				Phone:    value.Phone,
				Email:    value.Email,
				Address:  value.Address,
			}

			clientsPICUpdate = append(clientsPICUpdate, clientPICUpdate)
		} else {
			clientPICInsert := model.ClientPIC{
				IDClient: request.ID,
				Name:     value.Name,
				Phone:    value.Phone,
				Email:    value.Email,
				Address:  value.Address,
			}

			clientsPICInsert = append(clientsPICInsert, clientPICInsert)
		}
	}

	svc.rpo.UpdateClient(ctx, tx, client)
	svc.rpo.UpdateClientPIC(ctx, tx, &clientsPICUpdate)

	if len(clientsPICInsert) > 0 {
		svc.rpo.SaveClientPIC(ctx, tx, &clientsPICInsert)
	}

	return model.ClientResponse{
		ID:      client.ID,
		Name:    client.Name,
		Address: client.Address,
		Phone:   client.Phone,
	}
}

func (svc *Service) GetClients(ctx context.Context) []model.ClientResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	clients := svc.rpo.GetClients(ctx, tx)

	var result []model.ClientResponse
	for _, value := range *clients {
		client := model.ClientResponse{
			ID:      value.ID,
			Name:    value.Name,
			Address: value.Address,
			Phone:   value.Phone,
			Status:  value.Status,
		}

		result = append(result, client)
	}

	return result
}

func (svc *Service) GetClientsWithPagination(ctx context.Context, params *helper.PaginationParams) []model.ClientResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	clients := svc.rpo.GetClientsWithPagination(ctx, tx, params)

	var result []model.ClientResponse
	if len(*clients) > 0 {
		for _, value := range *clients {
			client := model.ClientResponse{
				ID:      value.ID,
				Name:    value.Name,
				Address: value.Address,
				Phone:   value.Phone,
			}

			result = append(result, client)
		}
	}

	return result
}

func (svc *Service) GetClient(ctx context.Context, id string) model.ClientResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	client, err := svc.rpo.GetClient(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result := model.ClientResponse{
		ID:        client.ID,
		Name:      client.Name,
		Address:   client.Address,
		Phone:     client.Phone,
		ClientPic: nil,
	}

	clientPic := svc.rpo.GetClientPIC(ctx, tx, id)

	var clientsPic []model.ClientPICResponse
	for _, value := range *clientPic {
		cp := model.ClientPICResponse{
			ID:      value.ID,
			Name:    value.Name,
			Phone:   value.Phone,
			Email:   value.Email,
			Address: value.Address,
		}

		clientsPic = append(clientsPic, cp)
	}

	result.ClientPic = clientsPic

	return result
}

func (svc *Service) DeleteClient(ctx context.Context, id string) {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer helper.CommitRollback(tx)

	_, errClient := svc.rpo.GetClient(ctx, tx, id)
	if errClient != nil {
		panic(exception.NewNotFoundError(errClient.Error()))
	}

	svc.rpo.DeleteClient(ctx, tx, id)
}

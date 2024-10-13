package project_manager

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/oktapascal/go-simpro/helper"
	"github.com/oktapascal/go-simpro/model"
	"github.com/oktapascal/go-simpro/web"
	"net/http"
)

type Handler struct {
	svc      model.ProjectManagerService
	validate *validator.Validate
}

func (hdl *Handler) SaveProjectManager() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(model.SaveProjectManagerRequest)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.StoreProjectManager(ctx, req)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) UpdateProjectManager() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(model.UpdateProjectManagerRequest)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.UpdateProjectManager(ctx, req)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) GetProjectManagersNoPagination() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		result := hdl.svc.GetProjectManagersNoPagination(ctx)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err := encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}

	}
}

func (hdl *Handler) GetOneProjectManager() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")

		ctx := request.Context()
		result := hdl.svc.GetOneProjectManager(ctx, id)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err := encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

func (hdl *Handler) DeleteProjectManager() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")

		ctx := request.Context()
		hdl.svc.DeleteProjectManager(ctx, id)

		svcResponse := web.DefaultResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   nil,
		}

		writer.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		err := encoder.Encode(svcResponse)
		if err != nil {
			panic(err)
		}
	}
}

package project

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
	svc      model.ProjectService
	validate *validator.Validate
}

func (hdl *Handler) SaveProject() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(model.SaveRequestProject)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.SaveProject(ctx, req)

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

func (hdl *Handler) UpdateProject() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(model.UpdateRequestProject)

		err := helper.DecodeRequest(request, req)
		if err != nil {
			panic(err)
		}

		err = hdl.validate.Struct(req)
		if err != nil {
			panic(err)
		}

		ctx := request.Context()
		result := hdl.svc.UpdateProject(ctx, req)

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

func (hdl *Handler) GetProjects() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		result := hdl.svc.GetProjects(ctx)

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

func (hdl *Handler) GetProject() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")

		ctx := request.Context()
		result := hdl.svc.GetProject(ctx, id)

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

func (hdl *Handler) SaveCloseProject() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (hdl *Handler) UpdateCloseProject() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (hdl *Handler) GetCloseProjects() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (hdl *Handler) GetCloseProject() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

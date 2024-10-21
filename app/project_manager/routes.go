package project_manager

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.ProjectManagerHandler
}

func (router *Router) InitializeRoute(route chi.Router) {
	route.Group(func(subroute chi.Router) {
		subroute.Use(middleware.AuthorizationCheckMiddleware)
		subroute.Use(middleware.VerifyAccessTokenMiddleware)
		subroute.Get("/project-managers/no-pagination", router.hdl.GetProjectManagersNoPagination())
		subroute.Get("/project-manager/{id}", router.hdl.GetOneProjectManager())
		subroute.Post("/project-manager", router.hdl.SaveProjectManager())
		subroute.Put("/project-manager/{id}", router.hdl.UpdateProjectManager())
		subroute.Delete("/project-manager/{id}", router.hdl.DeleteProjectManager())
	})
}

package project_manager

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.ProjectManagerHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api/project-manager", func(route chi.Router) {
		route.Use(middleware.AuthorizationCheckMiddleware)
		route.Use(middleware.VerifyAccessTokenMiddleware)
		route.Get("/no-pagination", router.hdl.GetProjectManagersNoPagination())
		route.Get("/{id}", router.hdl.GetOneProjectManager())
		route.Post("/", router.hdl.SaveProjectManager())
		route.Put("/{id}", router.hdl.UpdateProjectManager())
		route.Delete("/{id}", router.hdl.DeleteProjectManager())
	})
}

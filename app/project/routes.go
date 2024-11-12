package project

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.ProjectHandler
}

func (router *Router) InitializeRoutes(route chi.Router) {
	route.Route("/project", func(subroute chi.Router) {
		subroute.Group(func(children chi.Router) {
			children.Use(middleware.AuthorizationCheckMiddleware)
			children.Use(middleware.VerifyAccessTokenMiddleware)
			children.Get("/", router.hdl.GetProjects())
			children.Get("/{id}", router.hdl.GetProject())
			children.Post("/", router.hdl.SaveProject())
			children.Put("/{id}", router.hdl.UpdateProject())
		})
	})
}

package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.ClientHandler
}

func (router *Router) InitializeRoutes(route chi.Router) {
	route.Route("/client", func(subroute chi.Router) {
		subroute.Group(func(children chi.Router) {
			children.Use(middleware.AuthorizationCheckMiddleware)
			children.Use(middleware.VerifyAccessTokenMiddleware)
			children.Get("//with-pagination", router.hdl.GetClientsWithPagination())
			children.Get("/", router.hdl.GetClients())
			children.Get("/{id}", router.hdl.GetClient())
			children.Post("/", router.hdl.SaveClient())
			children.Put("/{id}", router.hdl.UpdateClient())
			children.Delete("/{id}", router.hdl.DeleteClient())
		})
	})
}

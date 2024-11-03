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
			children.Get("/clients/with-pagination", router.hdl.GetClientsWithPagination())
			children.Get("/clients", router.hdl.GetClients())
			children.Get("/client/{id}", router.hdl.GetClient())
			children.Post("/client", router.hdl.SaveClient())
			children.Put("/client/{id}", router.hdl.UpdateClient())
			children.Delete("/client/{id}", router.hdl.DeleteClient())
		})
	})
}

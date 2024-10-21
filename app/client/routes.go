package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.ClientHandler
}

func (router *Router) InitializeRoute(route chi.Router) {
	route.Group(func(subroute chi.Router) {
		subroute.Use(middleware.VerifyAccessTokenMiddleware)
		subroute.Use(middleware.AuthorizationCheckMiddleware)
		subroute.Get("/clients/pagination", router.hdl.GetAllClients())
		subroute.Get("/clients/no-pagination", router.hdl.GetClientsNoPagination())
		subroute.Get("/client/{id}", router.hdl.GetOneClient())
		subroute.Post("/client", router.hdl.SaveClient())
		subroute.Put("/client/{id}", router.hdl.UpdateClient())
		subroute.Delete("/client/{id}", router.hdl.DeleteClient())
	})
}

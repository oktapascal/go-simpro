package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.ClientHandler
}

func (router *Router) InitializeRoute(mux *chi.Mux) {
	mux.Route("/api", func(route chi.Router) {
		route.Use(middleware.AuthorizationCheckMiddleware)
		route.Use(middleware.VerifyAccessTokenMiddleware)
		route.Get("/clients/pagination", router.hdl.GetAllClients())
		route.Get("/clients/no-pagination", router.hdl.GetClientsNoPagination())
		route.Get("/client/{id}", router.hdl.GetOneClient())
		route.Post("/client", router.hdl.SaveClient())
		route.Put("/client/{id}", router.hdl.UpdateClient())
		route.Delete("/client/{id}", router.hdl.DeleteClient())
	})
}

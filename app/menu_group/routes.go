package menu_group

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.MenuGroupHandler
}

func (router *Router) InitializeRoute(route chi.Router) {
	route.Group(func(subroute chi.Router) {
		subroute.Use(middleware.AuthorizationCheckMiddleware)
		subroute.Use(middleware.VerifyAccessTokenMiddleware)
		subroute.Get("/menu-groups", router.hdl.GetAllMenuGroups())
		subroute.Get("/menu-group/{id}", router.hdl.GetOneMenuGroup())
		subroute.Post("/menu-group", router.hdl.SaveMenuGroup())
		subroute.Put("/menu-group/{id}", router.hdl.UpdateMenuGroup())
		subroute.Delete("/menu-group/{id}", router.hdl.DeleteMenuGroup())
	})
}

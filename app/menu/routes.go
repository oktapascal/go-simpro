package menu

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.MenuHandler
}

func (router *Router) InitializeRoute(route chi.Router) {
	route.Group(func(subroute chi.Router) {
		subroute.Use(middleware.AuthorizationCheckMiddleware)
		subroute.Use(middleware.VerifyAccessTokenMiddleware)
		subroute.Get("/menus", router.hdl.GetMenu())
	})
}

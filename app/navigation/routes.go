package navigation

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.NavigationHandler
}

func (router *Router) InitializeRoutes(route chi.Router) {
	route.Route("/navigation", func(subroute chi.Router) {
		subroute.Group(func(children chi.Router) {
			children.Use(middleware.AuthorizationCheckMiddleware)
			children.Use(middleware.VerifyAccessTokenMiddleware)
			children.Get("/", router.hdl.GetNavigation())
		})
	})
}

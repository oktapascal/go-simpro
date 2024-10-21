package login

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.LoginSessionHandler
}

func (router *Router) InitializeRoute(route chi.Router) {
	route.Route("/auth", func(routes chi.Router) {
		routes.Post("/login", router.hdl.Login())

		routes.Group(func(subroute chi.Router) {
			subroute.Use(middleware.AuthorizationCheckMiddleware)
			subroute.Use(middleware.VerifyRefreshTokenMiddleware)
			subroute.Get("/access-token", router.hdl.GetAccessToken())
		})

		routes.Group(func(subroute chi.Router) {
			subroute.Use(middleware.AuthorizationCheckMiddleware)
			subroute.Use(middleware.VerifyAccessTokenMiddleware)
			subroute.Post("/logout", router.hdl.Logout())
		})
	})
}

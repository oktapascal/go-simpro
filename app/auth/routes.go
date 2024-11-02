package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.AuthSessionHandler
}

func (router *Router) InitializeRoutes(route chi.Router) {
	route.Route("/auth", func(subroute chi.Router) {
		subroute.Post("/login", router.hdl.Login())

		subroute.Group(func(children chi.Router) {
			children.Use(middleware.AuthorizationCheckMiddleware)
			children.Use(middleware.VerifyRefreshTokenMiddleware)
			children.Get("/access-token", router.hdl.GetAccessToken())
		})

		subroute.Group(func(children chi.Router) {
			children.Use(middleware.AuthorizationCheckMiddleware)
			children.Use(middleware.VerifyAccessTokenMiddleware)
			children.Post("/logout", router.hdl.Logout())
		})
	})
}

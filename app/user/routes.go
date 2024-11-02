package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.UserHandler
}

func (router *Router) InitializeRoutes(route chi.Router) {
	route.Route("/user", func(subroute chi.Router) {
		subroute.Post("/", router.hdl.SaveUser())

		subroute.Group(func(children chi.Router) {
			children.Use(middleware.AuthorizationCheckMiddleware)
			children.Use(middleware.VerifyAccessTokenMiddleware)
			children.Post("/upload-photo", router.hdl.UpdateProfilePhotoUser())
		})
	})
}

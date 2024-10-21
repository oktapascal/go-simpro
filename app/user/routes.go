package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.UserHandler
}

func (router *Router) InitializeRoute(route chi.Router) {
	route.Route("/user", func(subroute chi.Router) {
		subroute.Get("/{username}/profile-photo", router.hdl.GetPhotoProfile())

		subroute.Group(func(subnestroute chi.Router) {
			subnestroute.Use(middleware.AuthorizationCheckMiddleware)
			subnestroute.Use(middleware.VerifyAccessTokenMiddleware)
			subnestroute.Get("/with-token", router.hdl.GetUserByToken())
			subnestroute.Post("/", router.hdl.SaveUser())
			subnestroute.Put("/", router.hdl.EditUser())
			subnestroute.Post("/upload-photo", router.hdl.UploadPhotoProfile())
		})
	})
}

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
		subroute.Group(func(children chi.Router) {
			children.Use(middleware.AuthorizationCheckMiddleware)
			children.Use(middleware.VerifyAccessTokenMiddleware)
			children.Group(func(subchildren chi.Router) {
				subchildren.Use(middleware.VerifyRootUserMiddleware)
				subchildren.Post("/", router.hdl.SaveUser())
			})
			children.Get("/retrieved-photo", router.hdl.GetUserPhotoProfile())
			children.Get("/with-auth", router.hdl.GetUserByToken())
			children.Post("/upload-photo", router.hdl.UpdateProfilePhotoUser())
			children.Put("/", router.hdl.UpdateUser())
		})
	})
}

package pic

import (
	"github.com/go-chi/chi/v5"
	"github.com/oktapascal/go-simpro/middleware"
	"github.com/oktapascal/go-simpro/model"
)

type Router struct {
	hdl model.PICHandler
}

func (router *Router) InitializeRoutes(route chi.Router) {
	route.Route("/pic", func(subroute chi.Router) {
		subroute.Group(func(children chi.Router) {
			children.Use(middleware.AuthorizationCheckMiddleware)
			children.Use(middleware.VerifyAccessTokenMiddleware)
			children.Get("//with-pagination", router.hdl.GetPICsWithPagination())
			children.Get("/", router.hdl.GetPICs())
			children.Get("/{id}", router.hdl.GetPIC())
			children.Post("/", router.hdl.SavePIC())
			children.Put("/{id}", router.hdl.UpdatePIC())
			children.Delete("/{id}", router.hdl.DeletePIC())
		})
	})
}

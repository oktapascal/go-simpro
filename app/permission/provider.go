package permission

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/oktapascal/go-simpro/model"
	"sync"
)

var (
	route     *Router
	routeOnce sync.Once

	hdl     *Handler
	hdlOnce sync.Once

	svc     *Service
	svcOnce sync.Once

	rpo     *Repository
	rpoOnce sync.Once

	ProviderSet = wire.NewSet(
		ProvideRoute,
		ProvideHandler,
		ProvideService,
		ProvideRepository,
		wire.Bind(new(model.PermissionHandler), new(*Handler)),
		wire.Bind(new(model.PermissionService), new(*Service)),
		wire.Bind(new(model.PermissionRepository), new(*Repository)),
	)
)

func ProvideRoute(hdl model.PermissionHandler) *Router {
	routeOnce.Do(func() {
		route = &Router{
			hdl: hdl,
		}
	})

	return route
}

func ProvideHandler(svc model.PermissionService, validate *validator.Validate) *Handler {
	hdlOnce.Do(func() {
		hdl = &Handler{
			svc:      svc,
			validate: validate,
		}
	})

	return hdl
}

func ProvideService(rpo model.PermissionRepository, db *sql.DB) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			rpo: rpo,
			db:  db,
		}
	})

	return svc
}

func ProvideRepository() *Repository {
	rpoOnce.Do(func() {
		rpo = new(Repository)
	})

	return rpo
}

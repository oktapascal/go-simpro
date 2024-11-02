package navigation

import (
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
		wire.Bind(new(model.NavigationHandler), new(*Handler)),
		wire.Bind(new(model.NavigationService), new(*Service)),
		wire.Bind(new(model.NavigationRepository), new(*Repository)),
	)
)

func ProvideRoute(hdl model.NavigationHandler) *Router {
	routeOnce.Do(func() {
		route = &Router{
			hdl: hdl,
		}
	})

	return route
}

func ProvideHandler(svc model.NavigationService) *Handler {
	hdlOnce.Do(func() {
		hdl = &Handler{
			svc: svc,
		}
	})

	return hdl
}

func ProvideService(rpo model.NavigationRepository) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			rpo: rpo,
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

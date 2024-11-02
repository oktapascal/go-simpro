//go:build wireinject
// +build wireinject

package navigation

import (
	"github.com/google/wire"
)

func Wire() *Router {
	panic(wire.Build(ProviderSet))
}

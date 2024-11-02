//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
)

func Wire() *Router {
	panic(wire.Build(ProviderSet))
}

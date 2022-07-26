// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/moecasts/kratos-layout-monorepo/internal/biz"
	"github.com/moecasts/kratos-layout-monorepo/internal/conf"
	"github.com/moecasts/kratos-layout-monorepo/internal/data"
	"github.com/moecasts/kratos-layout-monorepo/internal/server"
	"github.com/moecasts/kratos-layout-monorepo/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

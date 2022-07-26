package server

import (
	"net"
	"strings"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/moecasts/kratos-layout-monorepo/internal/conf"
	"github.com/soheilhy/cmux"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewMux, NewRegistarar)

func NewMux(conf *conf.Server) cmux.CMux {
	if !conf.Mux.Enable {
		return nil
	}
	lis, err := net.Listen(conf.Mux.Network, conf.Mux.Addr)
	if err != nil {
		panic(err)
	}
	m := cmux.New(lis)
	go func() {
		if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
			panic(err)
		}
	}()
	return m
}

func NewRegistarar(conf *conf.Registry, logger log.Logger) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme

	cli, err := consulAPI.NewClient(c)
	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}

	r := consul.New(cli)
	return r
}

package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/moecasts/kratos-layout-monorepo/internal/conf"
	"github.com/soheilhy/cmux"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(m cmux.CMux, c *conf.Server, logger log.Logger) *http.Server {
	if !c.Http.Enable {
		return nil
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Mux.Enable {
		opts = append(opts, http.Listener(m.Match(cmux.Any())))
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	return srv
}

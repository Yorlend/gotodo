package web

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHttpServer(
	lc fx.Lifecycle,
	lg *zap.SugaredLogger,
	r *mux.Router,
) *http.Server {

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			lg.With("addr", srv.Addr).Info("Starting HTTP server")

			go srv.Serve(ln)
			return nil
		},

		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

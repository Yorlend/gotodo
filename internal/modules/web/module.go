package web

import (
	"net/http"

	"github.com/yorlend/gotodo/internal/modules/web/handlers"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"web",
	fx.Provide(
		NewHttpServer,
		handlers.NewHomeHandler,
		handlers.NewTodosHandler,
		NewRouter,
	),
	fx.Invoke(func(*http.Server) {}),
)

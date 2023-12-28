package data

import "go.uber.org/fx"

var Module = fx.Module(
	"data",
	fx.Provide(
		NewDBClient,
		NewPgTodoRepo,
	),
)

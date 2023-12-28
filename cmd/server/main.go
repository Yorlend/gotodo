package main

import (
	"github.com/yorlend/gotodo/internal/modules/config"
	"github.com/yorlend/gotodo/internal/modules/data"
	"github.com/yorlend/gotodo/internal/modules/logger"
	"github.com/yorlend/gotodo/internal/modules/web"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		logger.Module,
		web.Module,
		config.Module,
		data.Module,
	)

	app.Run()
}

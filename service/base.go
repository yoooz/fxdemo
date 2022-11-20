package service

import "go.uber.org/fx"

func Modules() fx.Option {
	return fx.Provide(
		newEchoService,
	)
}

package handler

import (
	"net/http"

	"go.uber.org/fx"
)

type Route interface {
	http.Handler
	Pattern() string
}

func asRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func Modules() fx.Option {
	return fx.Provide(
		asRoute(newEchoHandler),
		asRoute(newHelloHandler),
	)
}

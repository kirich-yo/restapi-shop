package middleware

import (
	"net/http"
)

type Pair[F any, S any] struct {
	First F
	Second S
}

type Middleware func(next http.Handler, args ...interface{}) http.Handler

type MiddlewareWithArgs Pair[Middleware, []interface{}]

func Chain(middlewares ...MiddlewareWithArgs) Middleware {
	return func(next http.Handler, args ...interface{}) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i].First(next, middlewares[i].Second...)
		}

		return next
	}
}

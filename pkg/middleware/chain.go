package middleware

import "net/http"

type MiddleWare func(http.Handler) http.Handler

func Chain(middlerwares ...MiddleWare) MiddleWare {
	return func(next http.Handler) http.Handler {
		for i := len(middlerwares) - 1; i >= 0; i-- {
			next = middlerwares[i](next)
		}
		return next
	}
}
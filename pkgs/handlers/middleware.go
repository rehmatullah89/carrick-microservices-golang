package handlers

import "net/http"

func Middleware(httpHandler http.Handler, middleware ...func(h http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		httpHandler = mw(httpHandler)
	}

	return httpHandler
}

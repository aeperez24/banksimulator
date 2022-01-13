package middleware

import "net/http"

type HttpHandler func(w http.ResponseWriter, r *http.Request)
type Middleware interface {
	Filter(HttpHandler) HttpHandler
}



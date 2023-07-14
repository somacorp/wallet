package main

import (
	"fmt"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func ChainMiddleware(middlewares ...Middleware) Middleware {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler
	}
}

func (mw Middleware) With(middlewares ...Middleware) Middleware {
	return ChainMiddleware(append([]Middleware{mw}, middlewares...)...)
}

func (mw Middleware) Then(handler http.HandlerFunc) http.HandlerFunc {
	return mw(handler)
}

func Get(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions { // handle preflight
			w.WriteHeader(http.StatusNoContent)
			return
		} else if r.Method != http.MethodGet { // accept GET only
			http.Error(w, fmt.Sprintf("method %s is not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

func Post(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions { // handle preflight
			w.WriteHeader(http.StatusNoContent)
			return
		} else if r.Method != http.MethodPost { // accept POST only
			http.Error(w, fmt.Sprintf("method %s is not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

func (s *server) Guard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if !contains(s.conf.Origins, origin) {
			http.Error(w, fmt.Sprintf("forbidden origin %s", origin), http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

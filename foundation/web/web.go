// Package web asd dasklj asldk.
package web

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
	mw []MidFunc
}

func NewApp(mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

// HandleFunc THIS IS OUR FUNCTION.
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// WE CAN DO WHAT WE WANT BEFORE
		// INJECT FUNCTIONS

		handlerFunc(ctx, w, r)

		// INJECT FUNCTIONS
		// WE CAN DO WHAT WE WANT AFTER
	}

	a.ServeMux.HandleFunc(pattern, h)
}

// Package web asd dasklj asldk.
package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Logger func(ctx context.Context, msg string, args ...any)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) Encoder

type App struct {
	*http.ServeMux
	log Logger
	mw  []MidFunc
}

func NewApp(log Logger, mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		log:      log,
		mw:       mw,
	}
}

// HandleFunc THIS IS OUR FUNCTION.
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := setTraceID(r.Context(), uuid.New())

		// WE CAN DO WHAT WE WANT BEFORE
		// INJECT FUNCTIONS

		resp := handlerFunc(ctx, w, r)
		if err := Respond(ctx, w, resp); err != nil {
			a.log(ctx, "web", "ERROR", err)
			return
		}

		// INJECT FUNCTIONS
		// WE CAN DO WHAT WE WANT AFTER
	}

	a.ServeMux.HandleFunc(pattern, h)
}

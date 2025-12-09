// Package web asd dasklj asldk.
package web

import "net/http"

type App struct {
	*http.ServeMux
}

func NewApp() *App {
	return &App{
		ServeMux: http.NewServeMux(),
	}
}

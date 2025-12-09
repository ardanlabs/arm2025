// Package mux asdas da ads sa.
package mux

import (
	"net/http"

	"github.com/ardanlabs/service/app/domain/hackapp"
	"github.com/ardanlabs/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI() http.Handler {
	app := web.NewApp()

	hackapp.Routes(app)

	return app
}

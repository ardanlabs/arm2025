// Package mux asdas da ads sa.
package mux

import (
	"net/http"

	"github.com/ardanlabs/service/app/domain/hackapp"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger) http.Handler {
	app := web.NewApp(log.Info, mid.Logger(log))

	hackapp.Routes(app)

	return app
}

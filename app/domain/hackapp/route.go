package hackapp

import "github.com/ardanlabs/service/foundation/web"

func Routes(app *web.App) {
	app.HandleFunc("GET /hack", hack)
}

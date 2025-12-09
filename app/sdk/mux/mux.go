// Package mux asdas da ads sa.
package mux

import (
	"encoding/json"
	"net/http"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI() http.Handler {
	mux := http.NewServeMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string
		}{
			Status: "OK",
		}

		json.NewEncoder(w).Encode(status)
	}

	mux.HandleFunc("GET /hack", h)

	return mux
}

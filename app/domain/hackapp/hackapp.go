// Package hackapp asd aldkja lkds.
package hackapp

import (
	"encoding/json"
	"net/http"
)

func hack(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}

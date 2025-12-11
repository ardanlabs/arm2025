// Package hackapp asd aldkja lkds.
package hackapp

import (
	"context"
	"encoding/json"
	"net/http"
)

func hack(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return json.NewEncoder(w).Encode(status)
}

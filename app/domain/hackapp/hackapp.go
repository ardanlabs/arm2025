// Package hackapp asd aldkja lkds.
package hackapp

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func hack(ctx context.Context, w http.ResponseWriter, r *http.Request) web.Encoder {
	status := status{
		Status: "OK",
	}

	return status
}

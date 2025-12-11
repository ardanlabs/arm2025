// Package hackapp asd aldkja lkds.
package hackapp

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

func hack(ctx context.Context, w http.ResponseWriter, r *http.Request) web.Encoder {
	// DECODE INPUT
	// VALIDATE INPUT
	// BUSINESS LAYER
	// 	- RETURN ERROR
	//  - CONSTRUCT THE RESPONSE

	if n := rand.Intn(100); n%2 == 0 {
		return errs.Errorf(errs.InvalidArgument, "hack: %s[%s]", "param", "empty")
		//panic("OOOOOH NOOOOOO")
	}

	status := status{
		Status: "OK",
	}

	return status
}

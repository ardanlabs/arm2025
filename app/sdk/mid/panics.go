package mid

import (
	"context"
	"net/http"
	"runtime/debug"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

func Panics() web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (resp web.Encoder) {
			defer func() {
				if rec := recover(); rec != nil {
					trace := debug.Stack()
					resp = errs.Errorf(errs.InternalOnlyLog, "PANIC [%v] TRACE[%s]", rec, string(trace))
				}
			}()

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}

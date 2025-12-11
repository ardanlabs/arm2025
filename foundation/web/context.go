package web

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey int

const traceID ctxKey = 1

func setTraceID(ctx context.Context, uuid uuid.UUID) context.Context {
	return context.WithValue(ctx, traceID, uuid.String())
}

func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(traceID).(string)
	if !ok {
		return uuid.UUID{}.String()
	}

	return v
}

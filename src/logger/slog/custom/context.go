package custom

import (
	"context"

	"github.com/google/uuid"
)

var ctxKey = struct{}{}

func SetRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey, uuid.New().String())
}

func getRequestID(ctx context.Context) string {
	ctxValue, ok := ctx.Value(ctxKey).(string)
	if !ok {
		return "nothing in context"
	}

	return ctxValue
}

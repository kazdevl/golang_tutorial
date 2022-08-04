package rcontext

import (
	"context"
)

type contextKey struct{}

var rcontextKey = contextKey{}

type RequestContext struct {
	RequestID string
}

func Extract(ctx context.Context) *RequestContext {
	v := ctx.Value(rcontextKey)
	if v == nil {
		return nil
	}
	return v.(*RequestContext)
}

func Set(ctx context.Context) context.Context {
	return context.WithValue(ctx, rcontextKey, &RequestContext{RequestID: "string"})
}

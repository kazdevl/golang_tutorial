package econtext

import (
	"context"
	"errors"
)

type contextKey struct{}

var econtextKey = contextKey{}

type ErrorContext struct {
	Error error
}

func Extract(ctx context.Context) *ErrorContext {
	v := ctx.Value(econtextKey)
	if v == nil {
		return nil
	}
	return v.(*ErrorContext)
}

func Set(ctx context.Context) context.Context {
	return context.WithValue(ctx, econtextKey, &ErrorContext{Error: errors.New("sample")})
}

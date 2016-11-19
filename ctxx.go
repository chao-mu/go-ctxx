package ctxx

import (
	"context"
	"net/http"
)

// ContextFunc is meant to derive a new context.Context from an old, modifying
// as needed.
type ContextFunc func(ctx context.Context) context.Context

// Chain combines ContextFuncs, evaluating and passing along the
// context from left to right.
func Chain(links ...ContextFunc) ContextFunc {
	return func(ctx context.Context) context.Context {
		for _, link := range links {
			ctx = link(ctx)
		}

		return ctx
	}
}

// WithValue produces a ContextFunc that wraps context.WithValue.
// This is useful when combined with Chain (passing a bunch of WithValue
// results) and Wrap/WrapFunc to set several context values at once.
func WithValue(key, val interface{}) ContextFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, key, val)
	}
}

// RequestWithValue produces a new request whose context now includes
// key/val.
func RequestWithValue(req *http.Request, key, val interface{}) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), key, val))
}

func IsSet(ctx context.Context, key interface{}) bool {
	return ctx.Value(key) != nil
}

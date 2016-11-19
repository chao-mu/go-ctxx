package ctxx

import (
	"context"
	"database/sql"
)

func String(ctx context.Context, key interface{}) string {
	v := ctx.Value(key)
	if v == nil {
		return ""
	}

	switch v := v.(type) {
	case string:
		return v
	default:
		return ""
	}
}

func Uint(ctx context.Context, key interface{}) uint {
	v := ctx.Value(key)
	if v == nil {
		return 0
	}

	switch v := v.(type) {
	case uint:
		return v
	default:
		return 0
	}
}

func Int(ctx context.Context, key interface{}) int {
	v := ctx.Value(key)
	if v == nil {
		return 0
	}

	switch v := v.(type) {
	case int:
		return v
	default:
		return 0
	}
}

func Byte(ctx context.Context, key interface{}) byte {
	v := ctx.Value(key)
	if v == nil {
		return 0
	}

	switch v := v.(type) {
	case byte:
		return v
	default:
		return 0
	}
}

func Bool(ctx context.Context, key interface{}) bool {
	v := ctx.Value(key)
	if v == nil {
		return false
	}

	switch v := v.(type) {
	case bool:
		return v
	default:
		return false
	}
}

func Tx(ctx context.Context, key interface{}) *sql.Tx {
	v := ctx.Value(key)
	if v == nil {
		return nil
	}

	switch v := v.(type) {
	case *sql.Tx:
		return v
	default:
		return nil
	}
}

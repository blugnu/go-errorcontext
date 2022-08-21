package errorcontext

import (
	"context"
	"fmt"
)

// New creates a new ErrorWithContext, capturing the specified context and
// a string for the error.
//
// The error returned has a string representation of: "<text>"
func New(ctx context.Context, text string) error {
	return ErrorWithContext{ctx: ctx, string: text}
}

// Newf creates a new ErrorWithContext, capturing the specified context,
// a string and args to be substituted for any format specifiers in the
// string.
//
// The error returned has a string representation of: "<text>"
func Newf(ctx context.Context, text string, args ...interface{}) error {
	return ErrorWithContext{ctx: ctx, string: fmt.Sprintf(text, args...)}
}

// Wrap wraps a specified error, capturing the specified context and
// a string to use when representing the wrapped error.
//
// The error returned has a string representation of: "<text>: <err>"
func Wrap(ctx context.Context, err error, text string) error {
	return ErrorWithContext{ctx: ctx, error: err, string: text}
}

// Wrapf wraps a specified error, capturing the specified context a
// string to use when representing the wrapped error and args to be
// substituted for any format specifiers in the string.
//
// The error returned has a string representation of: "<text>: <err>"
func Wrapf(ctx context.Context, err error, text string, args ...interface{}) error {
	return ErrorWithContext{ctx: ctx, error: err, string: fmt.Sprintf(text, args...)}
}

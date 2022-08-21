package errorcontext

import (
	"context"
	"errors"
	"fmt"
)

// ErrorWithContext wraps an error, capturing a context and
// a string used to wrap the wrapped error in the string
// representation of the error.
type ErrorWithContext struct {
	ctx context.Context
	error
	string
}

// Error implements the error interface.
func (err ErrorWithContext) Error() string {
	if err.error != nil {
		return fmt.Sprintf("%s: %v", err.string, err.error)
	}
	return err.string
}

// Context returns the innermost context accessible from
// this error or any wrapped ErrorWithContext.
func (err ErrorWithContext) Context() context.Context {
	inner := ErrorWithContext{}
	if errors.As(err.error, &inner) {
		return inner.Context()
	}
	return err.ctx
}

// Unwrap implements unwrapping to return the wrapped error.
func (err ErrorWithContext) Unwrap() error {
	return err.error
}

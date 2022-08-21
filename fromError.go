package errorcontext

import (
	"context"
	"errors"
)

// FromError accepts a current context and an error and returns the context
// from the 'most wrapped' ErrorWithContext, or the supplied `context` if
// no ErrorWithContext is wrapped by the supplied `err`.
func FromError(ctx context.Context, err error) context.Context {
	errorWithContext := ErrorWithContext{}
	if errors.As(err, &errorWithContext) {
		return errorWithContext.Context()
	}
	return ctx
}

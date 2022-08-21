package errorcontext

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func Test_FromError_ReturnsDefaultContextIfNoErrorWithContextIsFound(t *testing.T) {

	// ARRANGE

	wanted := context.Background()
	err := errors.New("no context")

	// ACT

	got := FromError(wanted, err)

	// ASSERT

	if wanted != got {
		t.Errorf("wanted %v, got %v", wanted, got)
	}
}

func Test_FromError_ReturnsTheDeepestContext(t *testing.T) {

	// ARRANGE

	const value testkey = 1

	initial := context.Background()
	wanted := context.WithValue(initial, value, "value")

	err := ErrorWithContext{ctx: wanted, error: errors.New("no context")}
	innerwrap := fmt.Errorf("wrapped: %w", err)
	outerwrap := fmt.Errorf("wrapped: %w", innerwrap)

	// ACT

	got := FromError(initial, outerwrap)

	// ASSERT

	if wanted != got {
		t.Errorf("wanted %v, got %v", wanted, got)
	}
}

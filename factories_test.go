package errorcontext

import (
	"fmt"
	"context"
	"errors"
	"testing"
)

func Test_New(t *testing.T) {

	// ARRANGE
	ctx := context.Background()
	text := "error text"
	
	// ACT
	ewc := New(ctx, text).(ErrorWithContext)
	
	// ASSERT
	if (ewc.error != nil) {
		t.Error("unexpected wrapped error")
	}

	if (ewc.ctx != ctx) {
		t.Errorf("wanted %q, got %q", ctx, ewc.ctx)
	}

	if (ewc.string != text) {
		t.Errorf("wanted %q, got %q", text, ewc.string)
	}
}

func Test_Newf(t *testing.T) {

	// ARRANGE
	ctx := context.Background()
	text := "error text %q"
	arg := "with arg"
	
	// ACT
	ewc := Newf(ctx, text, arg).(ErrorWithContext)
	
	// ASSERT
	if (ewc.error != nil) {
		t.Error("unexpected wrapped error")
	}
	
	if (ewc.ctx != ctx) {
		t.Errorf("wanted %q, got %q", ctx, ewc.ctx)
	}
	
	wanted := fmt.Sprintf(text, arg)
	got := ewc.Error()
	if wanted != got {
		t.Errorf("wanted %q, got %q", wanted, got)
	}
}

func Test_Wrap(t *testing.T) {

	// ARRANGE
	ctx := context.Background()
	wrapped := errors.New("wrapped")
	text := "error text"
	
	// ACT
	ewc := Wrap(ctx, wrapped, text).(ErrorWithContext)
	
	// ASSERT
	if (ewc.error == nil) {
		t.Errorf("expected %q, got %q", wrapped, ewc.error)
	}

	if (ewc.ctx != ctx) {
		t.Errorf("wanted %q, got %q", ctx, ewc.ctx)
	}

	wanted := fmt.Sprintf("%s: %s", text, wrapped)
	got := ewc.Error()
	if wanted != got {
		t.Errorf("wanted %q, got %q", wanted, got)
	}
}

func Test_Wrapf(t *testing.T) {

	// ARRANGE
	ctx := context.Background()
	wrapped := errors.New("wrapped")
	text := "error text with %q"
	arg := "arg"
	
	// ACT
	ewc := Wrapf(ctx, wrapped, text, arg).(ErrorWithContext)
	
	// ASSERT
	if (ewc.error == nil) {
		t.Errorf("expected %q, got %q", wrapped, ewc.error)
	}

	if (ewc.ctx != ctx) {
		t.Errorf("wanted %q, got %q", ctx, ewc.ctx)
	}

	wanted := fmt.Sprintf("%s: %s", fmt.Sprintf(text, arg), wrapped)
	got := ewc.Error()
	if wanted != got {
		t.Errorf("wanted %q, got %q", wanted, got)
	}
}

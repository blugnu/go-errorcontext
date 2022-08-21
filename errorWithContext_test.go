package errorcontext

import (
	"context"
	"errors"
	"testing"
)

type testkey int

func TestErrorWithContextScenario(t *testing.T) {
	// ARRANGE
	const (
		keyHttp testkey = iota
		keyService
		keyDB
	)

	// Simulating a scenario where an api request is received via
	// http endpoint which adds http info to the context...
	ctx := context.Background()
	ctxApi := context.WithValue(ctx, keyHttp, struct {
		method string
		uri    string
	}{
		method: "GET",
		uri:    "/api/v1/document",
	})
	// .. then calls some service via mediator which just passes the context
	// on to the service...
	ctxService := ctxApi
	// .. which uses a repository to get a requested document from
	// a database, so adds db connection info to the context...
	ctxRepository := context.WithValue(ctxService, keyDB, struct {
		host string
		db   string
	}{
		host: "db2.prod",
		db:   "service-db",
	})
	// .. then gets an error when trying to SELECT the document data...
	dbErr := errors.New("table does not exist")
	// .. which the repository wraps with the repository context before
	// returning the error to the service...
	repositoryError := Wrap(ctxRepository, dbErr, "SELECT")
	// .. which wraps the error with service context before returning
	// it to the mediator...
	serviceError := Wrap(ctxService, repositoryError, "repository.getDocument")
	// .. which wraps the error with with the caller (api) context and
	// the handler signature before returning it to the caller (the api endpoint)...
	err := Wrap(ctxApi, serviceError, "*getDocument.Handler[getDocument.Request, getDocument.Result]")

	// .. which then processes the error..

	// ACT
	// if err := mediator.Perform(.. get document request ..); err != nil
	{
		ctx := FromError(ctx, err)
		http := ctx.Value(keyHttp)
		if http == nil {
			t.Error("no http in context")
		}

		// Service didn't add anything to the context!
		service := ctx.Value(keyService)
		if service != nil {
			t.Error("unexpected service in context")
		}

		db := ctx.Value(keyDB)
		if db == nil {
			t.Error("no database in context")
		}
	}

	ewc := ErrorWithContext{}
	ok := errors.As(err, &ewc)
	if !ok {
		t.Error("did not get ErrorWithContext")
	}
}

func TestThatErrorWithContextErrorFormatsCorrectly(t *testing.T) {
	// ARRANGE
	err := Wrap(context.Background(), errors.New("error"), "wraps")
	wanted := "wraps: error"

	// ACT
	got := err.Error()

	// ASSERT
	if wanted != got {
		t.Errorf("wanted %v, got %v", wanted, got)
	}
}

func TestThatErrorWithContextErrorUnwrapsCorrectly(t *testing.T) {
	// ARRANGE
	wanted := errors.New("error")
	err := Wrap(context.Background(), wanted, "wraps")

	// ACT
	got := errors.Unwrap(err)

	// ASSERT
	if wanted != got {
		t.Errorf("wanted %v, got %v", wanted, got)
	}
}

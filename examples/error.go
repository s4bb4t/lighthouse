package examples

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"github.com/s4bb4t/lighthouse/pkg/logger"
	"net/http"
)

// ExampleRouter demonstrates how to handle propagated SPError instances in the top-level handler layer.
//
// This pattern assumes that lower-level logic (service/app/repo layers) already constructs meaningful
// structured errors using sp.New(...) or sp.Wrap(...).
//
// You can switch on the HTTP status code to determine how to respond to clients. This gives you fine-grained
// control over response messages while preserving full error context.
func ExampleRouter() error {
	if err := ExampleApi(nil); err != nil {
		switch sp.Ensure(err).Code() {
		case http.StatusBadRequest, http.StatusNotFound:
			// These errors are user-facing and safe to return directly.
			return err

		case http.StatusInternalServerError:
			// For internal server errors, you may want to trim the error stack
			// to a user-safe level using .Spin(LevelError)
			return sp.Ensure(err).Spin(levels.LevelError)

		default:
			// Unknown/unclassified error — wrap it into a safe, generic 500.
			return sp.Wrap(sp.Ensure(err), sp.New(sp.Sample{
				Messages: map[string]string{
					sp.En: "Internal error",
				},
				Desc:     "Internal server error. We are working on it",
				Hint:     "Try again later",
				HttpCode: 500,
				Level:    levels.LevelUser,
			}))
		}
	}
	return nil
}

// ExampleApi represents the HTTP/business logic layer.
//
// It checks input, performs validations, and delegates to lower-level services.
// Errors are constructed or re-wrapped to provide appropriate HTTP semantics.
func ExampleApi(a any) error {
	if !exampleCheck(a) {
		return sp.New(sp.Sample{
			Messages: map[string]string{
				sp.En: "Bad request",
			},
			Desc:     "Invalid parameter: A",
			Hint:     "Fix input and try again",
			HttpCode: 400,
			Level:    levels.LevelUser,
		})
	}

	b, err := ExampleApp(a)
	if err != nil {
		// Known error — propagate as-is
		if sp.Ensure(err).Code() == http.StatusNotFound {
			return err
		}
		// Unknown error — wrap with high-level context
		return sp.Wrap(sp.Ensure(err), sp.New(sp.Sample{
			Messages: map[string]string{
				sp.En: "Internal error",
			},
			Desc:     "Unexpected error occurred in service layer",
			Hint:     "Try again later",
			HttpCode: 500,
			Level:    levels.LevelUser,
		}))
	}

	if b == nil {
		return sp.New(sp.Sample{
			Messages: map[string]string{
				sp.En: "Not Found",
			},
			Desc:     "No results found for the given input",
			Hint:     "Check if the ID or filter is valid",
			HttpCode: 404,
			Level:    levels.LevelUser,
		})
	}

	return nil
}

// ExampleApp represents the business layer that depends on internal helpers.
//
// It is expected that internal functions return already-wrapped SPError instances.
func ExampleApp(a any) (any, error) {
	_, err := exampleApp(a)
	if err != nil {
		logger.Noop().Error(sp.Ensure(err), levels.LevelError)
		return nil, err
	}
	return nil, nil
}

// exampleApp is an internal helper representing lower-level business logic.
//
// It wraps repository errors with additional service-level context.
func exampleApp(a any) (any, error) {
	_, err := ExampleRepo(nil)
	if err != nil {
		return nil, sp.Wrap(sp.Ensure(err), sp.New(sp.Sample{
			Messages: map[string]string{
				sp.En: "Db failed",
			},
			Desc:  "Failed to retrieve data from repository",
			Hint:  "Check repository implementation or DB state",
			Level: levels.LevelError,
		}))
	}
	return nil, nil
}

// ExampleRepo represents the data layer. It constructs base SPError values directly.
func ExampleRepo(a any) (any, error) {
	// Simulated no-result case — directly return 404
	if errors.Is(nil, sql.ErrNoRows) {
		return nil, sp.New(sp.Sample{
			Messages: map[string]string{
				sp.En: "Not found",
			},
			Desc:     "No matching records",
			Hint:     "Check query parameters",
			HttpCode: 404,
			Level:    levels.LevelUser,
		})
	}

	// Simulated DB error case — return wrapped internal error
	return nil, sp.New(sp.Sample{
		Messages: map[string]string{
			sp.En: "DB error",
		},
		Desc:     "Query execution failed",
		Hint:     "Inspect DB credentials or connection",
		HttpCode: 500,
		Level:    levels.LevelDebug,
		Cause:    fmt.Errorf("db closed"),
	})
}

// exampleCheck simulates parameter validation.
func exampleCheck(any) bool {
	return false
}

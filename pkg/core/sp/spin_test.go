package sp

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"testing"
)

// TODO: refactor tests

func TestSPError_Spin(t *testing.T) {
	root := Api()

	err := root.Spin(levels.LevelError)
	fmt.Println(err.Source())
}

func Api() *Error {
	err := App()
	return Wrap(err, Sample{
		Messages: map[string]string{
			En: "Internal",
		},
		Desc:  "Internal Error",
		Hint:  "Try Again later",
		Level: levels.LevelInfo,
	})
}

func App() *Error {
	err := DB()
	return Wrap(err, Sample{
		Messages: map[string]string{
			En: "App err",
		},
		Desc:  "Database error",
		Hint:  "Check repo layer",
		Level: levels.LevelError,
	})
}

func DB() *Error {
	return New(Sample{
		Messages: map[string]string{
			En: "Db connection failed",
		},
		Desc:  "Failed to connect to storage",
		Hint:  "check connection string, credentials, etc.",
		Level: levels.LevelDebug,
	})
}

func TestWrap(t *testing.T) {
	spFirst := New(Sample{
		Messages: map[string]string{
			En: "Error message",
		},
		Desc: "First Error",
		Hint: "Delete system32",
	})

	t.Log(Wrap(spFirst, Sample{
		Messages: map[string]string{
			En: "error message",
		},
		Desc: "Second Error",
		Hint: "read spFirst",
	}).Hint())
}

func TestWrapMethod(t *testing.T) {
	spFirst := New(Sample{
		Messages: map[string]string{
			En: "Error message",
		},
		Desc: "First Error",
		Hint: "Delete system32",
	})
	spSecond := New(Sample{
		Messages: map[string]string{
			En: "error message",
		},
		Desc: "Second Error",
		Hint: "read spFirst",
	})

	_ = spSecond.Wrap(spFirst)
	t.Log("ok")
}

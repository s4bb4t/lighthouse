package sp

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"testing"
)

func TestSPError_Spin(t *testing.T) {
	root := Api()

	err := root.Spin(levels.LevelMediumDebug)
	fmt.Println(err.ReadSource())

	b, _ := err.MarshalJSON()

	fmt.Println(string(b))
}

func Api() *Error {
	err := App()
	return Wrap(err, Err{
		Messages: map[string]string{
			En: "Internal",
		},
		Desc:  "Internal Error",
		Hint:  "Try Again later",
		Level: levels.LevelHighDebug,
	}).MustDone()
}

func App() *Error {
	err := DB()
	return Wrap(err, Err{
		Messages: map[string]string{
			En: "App err",
		},
		Desc:  "Database error",
		Hint:  "Check repo layer",
		Level: levels.LevelMediumDebug,
	}).MustDone()
}

func DB() *Error {
	return New(Err{
		Messages: map[string]string{
			En: "Db connection failed",
		},
		Desc:  "Failed to connect to storage",
		Hint:  "check connection string, credentials, etc.",
		Level: levels.LevelDeepDebug,
	}).MustDone()
}

func TestWrap(t *testing.T) {
	spFirst := New(Err{
		Messages: map[string]string{
			En: "Error message",
		},
		Desc: "First Error",
		Hint: "Delete system32",
	}).MustDone()

	t.Log(Wrap(spFirst, Err{
		Messages: map[string]string{
			En: "error message",
		},
		Desc: "Second Error",
		Hint: "read spFirst",
	}).ReadHint())
}

func TestWrapMethod(t *testing.T) {
	spFirst := New(Err{
		Messages: map[string]string{
			En: "Error message",
		},
		Desc: "First Error",
		Hint: "Delete system32",
	})
	spSecond := New(Err{
		Messages: map[string]string{
			En: "error message",
		},
		Desc: "Second Error",
		Hint: "read spFirst",
	})

	_ = spSecond.Wrap(spFirst)
	t.Log("ok")
}

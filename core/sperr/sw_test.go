package sperr

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/core/levels"
	"testing"
)

func TestSPError_Spin(t *testing.T) {
	root := Api()

	err := root.Spin(levels.LevelMediumDebug)
	fmt.Println(err)
	fmt.Println()
	err = root.Spin(levels.LevelDeepDebug)
	fmt.Println(err)
}

func Api() *SPError {
	err := App()
	return Wrap(err, Err{
		Messages: map[string]string{
			En: "Internal",
		},
		Desc:  "Internal Error",
		Hint:  "Try Again later",
		Path:  "api",
		Level: levels.LevelHighDebug,
	})
}
func App() *SPError {
	err := DB()
	return Wrap(err, Err{
		Messages: map[string]string{
			En: "App err",
		},
		Desc:  "Database error",
		Hint:  "Check repo layer",
		Path:  "app",
		Level: levels.LevelMediumDebug,
	})
}

func DB() *SPError {
	return SP(Err{
		Messages: map[string]string{
			En: "Db connection failed",
		},
		Desc:  "Failed to connect to database",
		Hint:  "check connection string, credentials, etc.",
		Path:  "Db",
		Level: levels.LevelDeepDebug,
	})
}

func TestWrap(t *testing.T) {
	spFirst := SP(Err{
		Messages: map[string]string{
			En: "Error message",
		},
		Desc: "First Error",
		Hint: "Delete system32",
		Path: "TestUsage()",
	})

	t.Log(Wrap(spFirst, Err{
		Messages: map[string]string{
			En: "error message",
		},
		Desc: "Second Error",
		Hint: "read spFirst",
		Path: "TestUsage",
	}).ReadHint())
}

func TestWrapMethod(t *testing.T) {
	spFirst := SP(Err{
		Messages: map[string]string{
			En: "Error message",
		},
		Desc: "First Error",
		Hint: "Delete system32",
		Path: "TestUsage()",
	})
	spSecond := SP(Err{
		Messages: map[string]string{
			En: "error message",
		},
		Desc: "Second Error",
		Hint: "read spFirst",
		Path: "TestUsage",
	})

	spSecond.Wrap(spFirst)
	t.Log("ok")
}

package sperr

import (
	"fmt"
	"testing"
)

func TestSPError_Spin(t *testing.T) {
	root := Api()

	err := root.Spin(LevelHighUser)
	fmt.Println(err)
	fmt.Println()
	fmt.Println()
	err = root.Spin(LevelMediumDebug)
	fmt.Println(err)
	fmt.Println()
	fmt.Println()
	err = root.Spin(LevelDeepDebug)
	fmt.Println(err)
}

func Api() *SPError {
	err := App()
	return Wrap(err, Fields{
		Messages: map[string]string{
			En: "Internal",
		},
		Desc:  "Internal Error",
		Hint:  "Try Again later",
		Path:  "api",
		Level: LevelHighDebug,
	})
}
func App() *SPError {
	err := DB()
	return Wrap(err, Fields{
		Messages: map[string]string{
			En: "App err",
		},
		Desc:  "Database error",
		Hint:  "Check repo layer",
		Path:  "app",
		Level: LevelMediumDebug,
	})
}

func DB() *SPError {
	return SP(Fields{
		Messages: map[string]string{
			En: "Db connection failed",
		},
		Desc:  "Failed to connect to database",
		Hint:  "check connection string, credentials, etc.",
		Path:  "Db",
		Level: LevelDeepDebug,
	})
}

func TestWrap(t *testing.T) {
	spFirst := SP(Fields{
		Messages: map[string]string{
			En: "Error message",
		},
		Desc: "First Error",
		Hint: "Delete system32",
		Path: "TestUsage()",
	})

	t.Log(Wrap(spFirst, Fields{
		Messages: map[string]string{
			En: "error message",
		},
		Desc: "Second Error",
		Hint: "read spFirst",
		Path: "TestUsage",
	}).ReadHint())
}

func TestWrapMethod(t *testing.T) {
	spFirst := SP(Fields{
		Messages: map[string]string{
			En: "Error message",
		},
		Desc: "First Error",
		Hint: "Delete system32",
		Path: "TestUsage()",
	})
	spSecond := SP(Fields{
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

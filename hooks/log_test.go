package hooks

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/core/levels"
	"github.com/s4bb4t/lighthouse/core/sp"
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	zap.NewNop().Error("test", Zap(Api(), levels.LevelMediumDebug)...)
	for _, v := range Zap(Api(), levels.LevelMediumDebug) {
		fmt.Println(v)
	}
}

func Api() *sp.Error {
	err := App()
	return sp.Wrap(err, sp.Err{
		Messages: map[string]string{
			sp.En: "Internal",
		},
		Desc:  "Internal Error",
		Hint:  "Try Again later",
		Path:  "api",
		Level: levels.LevelHighDebug,
	}).MustDone()
}
func App() *sp.Error {
	err := DB()
	return sp.Wrap(err, sp.Err{
		Messages: map[string]string{
			sp.En: "App err",
		},
		Desc:  "Database error",
		Hint:  "Check repo layer",
		Path:  "app",
		Level: levels.LevelMediumDebug,
	}).MustDone()
}

func DB() *sp.Error {
	return sp.SP(sp.Err{
		Messages: map[string]string{
			sp.En: "Db connection failed",
		},
		Desc:  "Failed to connect to database",
		Hint:  "check connection string, credentials, etc.",
		Path:  "Db",
		Level: levels.LevelDeepDebug,
	}).MustDone()
}

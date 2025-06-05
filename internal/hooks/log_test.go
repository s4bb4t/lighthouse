package hooks

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	sp2 "github.com/s4bb4t/lighthouse/pkg/core/sp"
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	zap.NewNop().Error("test", Zap(Api(), levels.LevelError)...)
	for _, v := range Zap(Api(), levels.LevelError) {
		fmt.Println(v)
	}
}

func Api() *sp2.Error {
	err := App()
	return sp2.Wrap(err, sp2.Sample{
		Messages: map[string]string{
			sp2.En: "Internal",
		},
		Desc:  "Internal Error",
		Hint:  "Try Again later",
		Level: levels.LevelInfo,
	})
}
func App() *sp2.Error {
	err := DB()
	return sp2.Wrap(err, sp2.Sample{
		Messages: map[string]string{
			sp2.En: "App err",
		},
		Desc:  "Database error",
		Hint:  "Check repo layer",
		Level: levels.LevelError,
	})
}

func DB() *sp2.Error {
	return sp2.New(sp2.Sample{
		Messages: map[string]string{
			sp2.En: "Db connection failed",
		},
		Desc:  "Failed to connect to storage",
		Hint:  "check connection string, credentials, etc.",
		Level: levels.LevelDebug,
	})
}

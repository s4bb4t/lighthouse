package telegram

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/internal/hooks"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"testing"
	"time"
)

func TestBot(t *testing.T) {
	b, err := New("7410446656:AAG2MSNlHI6PMejIxz4MGZd-nLaSKGnhNt0", nil)
	if err != nil {
		t.Fatal(err)
	}

	b.StartLocalWebHook(wh, "8081")
	time.Sleep(2 * time.Second)

	E := sp.New(sp.Err{
		Messages: map[string]string{
			"en": "failed to read user's id",
		},
		Desc:  "Invalid user id",
		Hint:  "Your user id is invalid - check what you tries to save",
		Level: levels.LevelError,
		Meta: map[string]any{
			"key":   "dads",
			"id":    "31221",
			"bytes": "123 1232 123 12 321 3",
		},
	}).MustDone()

	err = b.Error(E, "")
	if err != nil {
		fmt.Println(hooks.Slog(sp.Ensure(err), levels.LevelDebug))
	}
	select {}
}

package telegram

import (
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"hash"
)

var (
	sendErr hash.Hash
)

func init() {
	var err error
	sendErr, err = sp.Registry.Reg(sp.New(sp.Err{
		Messages: map[string]string{
			sp.En: "failed to send message",
		},
		Desc:  "failed to send message",
		Hint:  "Check underlying error",
		Level: levels.LevelError,
	}).MustDone())
	if err != nil {
		panic(err)
	}
}

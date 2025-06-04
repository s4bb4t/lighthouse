package main

import (
	"github.com/s4bb4t/lighthouse/pkg/api/lighthouse"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	sp2 "github.com/s4bb4t/lighthouse/pkg/core/sp"
	"github.com/s4bb4t/lighthouse/pkg/logger"
)

func main() {
	lh := lighthouse.New(logger.Dev, "7410446656:AAG2MSNlHI6PMejIxz4MGZd-nLaSKGnhNt0")
	lh.Debug("test")

	err := sp2.New(sp2.Err{
		Messages: map[string]string{
			sp2.En: "Db connection failed",
		},
		Desc:  "Failed to connect to storage",
		Hint:  "check connection string, credentials, etc.",
		Level: levels.LevelDeepDebug,
	}).MustDone()

	lh.Error(err, levels.LevelDeepDebug)
	select {}
}

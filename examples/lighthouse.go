package examples

import (
	"github.com/s4bb4t/lighthouse/api/lighthouse"
	"github.com/s4bb4t/lighthouse/core/sp"
	"github.com/s4bb4t/lighthouse/kibana"
	"github.com/s4bb4t/lighthouse/logger"
	"github.com/s4bb4t/lighthouse/tg"
	"os"
)

func main() {
	lh := lighthouse.ManualNew(tg.New(), kibana.NewHook(kibana.Config{
		URL:         "",
		IndexPrefix: "",
		BatchSize:   0,
		FlushPeriod: 0,
		APIKey:      "",
		Environment: "",
	}), logger.New(logger.Local, sp.En, os.Stdout))

	lh.Debug("test")
	lh.Fire()
}

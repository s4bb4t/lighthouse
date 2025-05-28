package kibana

import "time"

type Config struct {
	URL         string
	IndexPrefix string
	BatchSize   int
	FlushPeriod time.Duration
	APIKey      string
	Environment string
}

package kibana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/s4bb4t/lighthouse/core/sp"
)

type Client struct {
	config Config
	buffer [][]byte
	sync.Mutex
	client *http.Client
}

func NewClient(cfg Config) *Client {
	c := &Client{
		config: cfg,
		buffer: make([][]byte, 0, cfg.BatchSize),
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}

	go c.startFlushing()
	return c
}

func (c *Client) LogError(e *sp.Error) error {
	c.Lock()
	defer c.Unlock()

	b, err := e.MarshalJSON()
	if err != nil {
		return fmt.Errorf("marshaling error: %w", err)
	}

	c.buffer = append(c.buffer, b)

	if len(c.buffer) >= c.config.BatchSize {
		return c.flush()
	}

	return nil
}

func (c *Client) flush() error {
	if len(c.buffer) == 0 {
		return nil
	}

	req, err := http.NewRequest("POST", c.config.URL+"/_bulk", bytes.NewBuffer(c.createBulkPayload()))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-ndjson")
	req.Header.Set("Authorization", "ApiKey "+c.config.APIKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("elasticsearch error: status %d", resp.StatusCode)
	}

	c.buffer = c.buffer[:0]
	return nil
}

func (c *Client) startFlushing(wantErrors bool) chan error {
	errChan := make(chan error)

	if !wantErrors {
		go func() {
			ticker := time.NewTicker(c.config.FlushPeriod)
			for range ticker.C {
				c.Lock()
				_ = c.flush()
				c.Unlock()
			}
		}()
	}

	go func() {
		ticker := time.NewTicker(c.config.FlushPeriod)
		for range ticker.C {
			c.Lock()
			errChan <- c.flush()
			c.Unlock()
		}
	}()

	return errChan
}

func (c *Client) createBulkPayload() []byte {
	var payload []byte

	for _, log := range c.buffer {
		indexLine := struct {
			Index struct {
				Index string `json:"_index"`
			} `json:"index"`
		}{}

		indexLine.Index.Index = fmt.Sprintf("%s-%s",
			c.config.IndexPrefix,
			time.Now().Format("2006.01.02"))

		metaJSON, _ := json.Marshal(indexLine)
		payload = append(payload, metaJSON...)
		payload = append(payload, '\n')

		payload = append(payload, log...)
		payload = append(payload, '\n')
	}

	return payload
}

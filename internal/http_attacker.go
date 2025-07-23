package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type HTTPConfig struct {
	URL    string
	RPS    int
	Dur    time.Duration
	Method string
	Body   string
}

func RunHTTPTest(cfg HTTPConfig) error {
	var wg sync.WaitGroup
	var totalRequests uint64
	var failedRequests uint64
	latencies := make(chan time.Duration, cfg.RPS*10)

	ticker := time.NewTicker(time.Second / time.Duration(cfg.RPS))
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Dur)
	defer cancel()

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	fmt.Printf(" Attacking %s at %d RPS for %v\n\n", cfg.URL, cfg.RPS, cfg.Dur)

	startTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			goto Report
		case <-ticker.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				start := time.Now()

				req, _ := http.NewRequest(cfg.Method, cfg.URL, nil)
				if cfg.Body != "" {
					req, _ = http.NewRequest(cfg.Method, cfg.URL, bytes.NewBufferString(cfg.Body))
					req.Header.Set("Content-Type", "application/json")
				}

				resp, err := client.Do(req)
				if err != nil {
					atomic.AddUint64(&failedRequests, 1)
					return
				}

				io.Copy(io.Discard, resp.Body)
				resp.Body.Close()

				if resp.StatusCode >= 400 {
					atomic.AddUint64(&failedRequests, 1)
				}

				latencies <- time.Since(start)
				atomic.AddUint64(&totalRequests, 1)
			}()
		}
	}

Report:
	wg.Wait()
	close(latencies)
	duration := time.Since(startTime)

	printHTTPReport(totalRequests, failedRequests, latencies, duration, cfg)
	return nil
}

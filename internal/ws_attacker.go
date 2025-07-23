package internal

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

type WSConfig struct {
	URL         string
	Connections int
	Messages    int
	Interval    time.Duration
	MessageData string
}

type WSResult struct {
	TotalMessages uint64
	Errors        uint64
	MinLatency    time.Duration
	MaxLatency    time.Duration
	Latencies     []time.Duration
}

func RunWSTest(cfg WSConfig) error {
	var wg sync.WaitGroup
	result := WSResult{Latencies: make([]time.Duration, 0, cfg.Connections*cfg.Messages)}

	fmt.Printf(" Starting WebSocket test to %s\n", cfg.URL)
	fmt.Printf("• Connections: %d\n• Messages per conn: %d\n• Interval: %v\n\n", 
		cfg.Connections, cfg.Messages, cfg.Interval)

	startTime := time.Now()

	for i := 0; i < cfg.Connections; i++ {
		wg.Add(1)
		go func(connID int) {
			defer wg.Done()
			dialer := websocket.Dialer{
				HandshakeTimeout: 5 * time.Second,
			}

			conn, _, err := dialer.Dial(cfg.URL, nil)
			if err != nil {
				atomic.AddUint64(&result.Errors, 1)
				log.Printf("Connection %d failed: %v", connID, err)
				return
			}
			defer conn.Close()

			for msgID := 0; msgID < cfg.Messages; msgID++ {
				start := time.Now()

				// Kirim pesan
				if err := conn.WriteMessage(websocket.TextMessage, []byte(cfg.MessageData)); err != nil {
					atomic.AddUint64(&result.Errors, 1)
					continue
				}

				// Terima balasan
				_, _, err := conn.ReadMessage()
				if err != nil {
					atomic.AddUint64(&result.Errors, 1)
					continue
				}

				latency := time.Since(start)
				result.Latencies = append(result.Latencies, latency)
				atomic.AddUint64(&result.TotalMessages, 1)

				time.Sleep(cfg.Interval)
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(startTime)

	printWSReport(result, duration, cfg)
	return nil
}

func printWSReport(result WSResult, duration time.Duration, cfg WSConfig) {
	// [Output formatting and statistics calculation]
}

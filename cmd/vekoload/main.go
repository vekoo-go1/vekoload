package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/vekoload/internal"
)

func main() {
	app := &cli.App{
		Name:    "vekoload",
		Version: "1.0.0",
		Usage:   "High-performance load testing tool",
		Commands: []*cli.Command{
			{
				Name:  "http",
				Usage: "HTTP load testing",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "url", Required: true},
					&cli.IntFlag{Name: "rps", Value: 100},
					&cli.DurationFlag{Name: "duration", Value: 30 * time.Second},
					&cli.StringFlag{Name: "method", Value: "GET"},
					&cli.StringFlag{Name: "data"},
				},
				Action: func(c *cli.Context) error {
					config := internal.HTTPConfig{
						URL:    c.String("url"),
						RPS:    c.Int("rps"),
						Dur:    c.Duration("duration"),
						Method: c.String("method"),
						Body:   c.String("data"),
					}
					return internal.RunHTTPTest(config)
				},
			},
			{
				Name:  "ws",
				Usage: "WebSocket load testing",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "url", Required: true},
					&cli.IntFlag{Name: "connections", Value: 100},
					&cli.IntFlag{Name: "messages", Value: 10},
					&cli.DurationFlag{Name: "interval", Value: 1 * time.Second},
					&cli.StringFlag{Name: "data", Value: "ping"},
				},
				Action: func(c *cli.Context) error {
					config := internal.WSConfig{
						URL:         c.String("url"),
						Connections: c.Int("connections"),
						Messages:    c.Int("messages"),
						Interval:    c.Duration("interval"),
						MessageData: c.String("data"),
					}
					return internal.RunWSTest(config)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

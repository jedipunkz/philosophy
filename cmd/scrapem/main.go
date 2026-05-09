package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jedipunkz/philosophy/internal/config"
	"github.com/jedipunkz/philosophy/internal/scrapem"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("usage: scrapem <run|watch> --config scrapem.yaml")
	}

	switch os.Args[1] {
	case "run":
		fs := flag.NewFlagSet("run", flag.ExitOnError)
		configPath := fs.String("config", "scrapem.yaml", "path to scrapem.yaml")
		if err := fs.Parse(os.Args[2:]); err != nil {
			return err
		}
		cfg, err := config.Load(*configPath)
		if err != nil {
			return err
		}
		return scrapem.New(cfg).Run(context.Background())

	case "watch":
		fs := flag.NewFlagSet("watch", flag.ExitOnError)
		configPath := fs.String("config", "scrapem.yaml", "path to scrapem.yaml")
		if err := fs.Parse(os.Args[2:]); err != nil {
			return err
		}
		cfg, err := config.Load(*configPath)
		if err != nil {
			return err
		}
		interval, err := time.ParseDuration(cfg.Scrape.Interval)
		if err != nil {
			return fmt.Errorf("invalid scrape.interval: %w", err)
		}
		if interval <= 0 {
			return errors.New("scrape.interval must be positive")
		}

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		runner := scrapem.New(cfg)
		if err := runner.Run(ctx); err != nil {
			return err
		}

		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-ticker.C:
				if err := runner.Run(ctx); err != nil {
					log.Printf("scrape run failed: %v", err)
				}
			}
		}

	default:
		return fmt.Errorf("unknown command %q", os.Args[1])
	}
}

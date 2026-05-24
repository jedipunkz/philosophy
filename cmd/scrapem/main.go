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
		return errors.New("usage: scrapem <run|watch|dedupe> --config research-scrape.yaml")
	}

	switch os.Args[1] {
	case "dedupe":
		fs := flag.NewFlagSet("dedupe", flag.ExitOnError)
		configPath := fs.String("config", "research-scrape.yaml", "path to scrape config yaml")
		dryRun := fs.Bool("dry-run", false, "log actions without removing files")
		if err := fs.Parse(os.Args[2:]); err != nil {
			return err
		}
		cfg, err := config.Load(*configPath)
		if err != nil {
			return err
		}
		return scrapem.New(cfg).Dedupe(context.Background(), *dryRun)

	case "run":
		fs := flag.NewFlagSet("run", flag.ExitOnError)
		configPath := fs.String("config", "research-scrape.yaml", "path to scrape config yaml")
		maxDuration := fs.Duration("max-duration", 0, "stop the run cleanly after this duration")
		if err := fs.Parse(os.Args[2:]); err != nil {
			return err
		}
		cfg, err := config.Load(*configPath)
		if err != nil {
			return err
		}
		// Honor SIGTERM/SIGINT so an interrupted scheduled run exits 0 after
		// flushing seen state, letting the workflow's commit step still run.
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()
		if *maxDuration > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, *maxDuration)
			defer cancel()
		}
		if err := scrapem.New(cfg).Run(ctx); err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				log.Printf("scrape run interrupted: %v", err)
				return nil
			}
			return err
		}
		return nil

	case "watch":
		fs := flag.NewFlagSet("watch", flag.ExitOnError)
		configPath := fs.String("config", "research-scrape.yaml", "path to scrape config yaml")
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

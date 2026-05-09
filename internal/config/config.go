package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Vault    VaultConfig     `yaml:"vault"`
	Scrape   ScrapeConfig    `yaml:"scrape"`
	Sources  []SourceConfig  `yaml:"sources"`
	Keywords []KeywordConfig `yaml:"keywords"`
}

type VaultConfig struct {
	Root  string `yaml:"root"`
	Inbox string `yaml:"inbox"`
}

type ScrapeConfig struct {
	Interval        string `yaml:"interval"`
	UserAgent       string `yaml:"user_agent"`
	MaxResults      int    `yaml:"max_results"`
	MaxPDFChars     int    `yaml:"max_pdf_chars"`
	MaxPDFBytes     int64  `yaml:"max_pdf_bytes"`
	RequestTimeout  string `yaml:"request_timeout"`
	RequestDelay    string `yaml:"request_delay"`
	RefreshExisting bool   `yaml:"refresh_existing"`
}

type SourceConfig struct {
	Name      string `yaml:"name"`
	Enabled   bool   `yaml:"enabled"`
	Type      string `yaml:"type"`
	BaseURL   string `yaml:"base_url"`
	SearchURL string `yaml:"search_url"`
}

type KeywordConfig struct {
	Name    string   `yaml:"name"`
	Queries []string `yaml:"queries"`
	Sources []string `yaml:"sources"`
	Tags    []string `yaml:"tags"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	if cfg.Vault.Root == "" {
		cfg.Vault.Root = "."
	}
	if cfg.Vault.Inbox == "" {
		cfg.Vault.Inbox = "inbox"
	}
	if cfg.Scrape.Interval == "" {
		cfg.Scrape.Interval = "24h"
	}
	if cfg.Scrape.UserAgent == "" {
		cfg.Scrape.UserAgent = "scrapem/0.1"
	}
	if cfg.Scrape.MaxResults <= 0 {
		cfg.Scrape.MaxResults = 5
	}
	if cfg.Scrape.MaxPDFChars <= 0 {
		cfg.Scrape.MaxPDFChars = 30000
	}
	if cfg.Scrape.MaxPDFBytes <= 0 {
		cfg.Scrape.MaxPDFBytes = 20 * 1024 * 1024
	}
	if cfg.Scrape.RequestTimeout == "" {
		cfg.Scrape.RequestTimeout = "20s"
	}
	if cfg.Scrape.RequestDelay == "" {
		cfg.Scrape.RequestDelay = "2s"
	}
	if len(cfg.Sources) == 0 {
		return Config{}, fmt.Errorf("sources must not be empty")
	}
	if len(cfg.Keywords) == 0 {
		return Config{}, fmt.Errorf("keywords must not be empty")
	}
	for _, source := range cfg.Sources {
		if source.Name == "" {
			return Config{}, fmt.Errorf("source.name must not be empty")
		}
		if source.SearchURL == "" {
			return Config{}, fmt.Errorf("source %q requires search_url", source.Name)
		}
	}
	for _, keyword := range cfg.Keywords {
		if keyword.Name == "" {
			return Config{}, fmt.Errorf("keyword.name must not be empty")
		}
		if len(keyword.Queries) == 0 {
			return Config{}, fmt.Errorf("keyword %q requires at least one query", keyword.Name)
		}
		if len(keyword.Sources) == 0 {
			return Config{}, fmt.Errorf("keyword %q requires at least one source", keyword.Name)
		}
	}

	return cfg, nil
}

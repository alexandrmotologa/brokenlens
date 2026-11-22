package config

import (
	"testing"
)

func TestParseFlagsVersion(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"LongFlag", []string{"--version"}},
		{"ShortFlag", []string{"-v"}},
		{"Subcommand", []string{"version"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := ParseFlags(tc.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !cfg.ShowVersion {
				t.Errorf("expected ShowVersion to be true for args %v", tc.args)
			}
		})
	}
}

func TestParseFlagsTargetURL(t *testing.T) {
	cfg, err := ParseFlags([]string{"https://example.com", "--concurrency", "10", "--max-depth", "3"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.TargetURL != "https://example.com" {
		t.Errorf("expected target URL https://example.com, got %s", cfg.TargetURL)
	}
	if cfg.Concurrency != 10 {
		t.Errorf("expected concurrency 10, got %d", cfg.Concurrency)
	}
	if cfg.MaxDepth != 3 {
		t.Errorf("expected max depth 3, got %d", cfg.MaxDepth)
	}
}

func TestParseFlagsMissingTarget(t *testing.T) {
	_, err := ParseFlags([]string{"--concurrency", "10"})
	if err == nil {
		t.Fatal("expected error when target URL is missing, got nil")
	}
}

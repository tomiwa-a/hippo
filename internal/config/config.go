package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	WatchPaths []string  `mapstructure:"watch"`
	Ignore     []string  `mapstructure:"ignore"`
	DBPath     string    `mapstructure:"db_path"`
	MaxSize    int64     `mapstructure:"max_size"`
	Workers    int       `mapstructure:"workers"`
	Embedding  Embedding `mapstructure:"embedding"`
}

type Embedding struct {
	Provider string `mapstructure:"provider"` // ollama, openai, etc.
	BaseURL  string `mapstructure:"base_url"`
	Model    string `mapstructure:"model"`
}

func Load() (*Config, error) {
	viper.SetConfigName("hippo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(os.Getenv("HOME") + "/.hippo")

	viper.SetDefault("watch", []string{"."})
	viper.SetDefault("ignore", []string{".git", "node_modules", "dist", "vendor", "ui"})
	viper.SetDefault("db_path", "hippo.db")
	viper.SetDefault("max_size", 10*1024*1024) // 10MB default
	viper.SetDefault("workers", runtime.NumCPU())
	viper.SetDefault("embedding.provider", "ollama")
	viper.SetDefault("embedding.base_url", "http://localhost:11434")
	viper.SetDefault("embedding.model", "nomic-embed-text")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	// Resolve absolute paths for WatchPaths
	for i, path := range cfg.WatchPaths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve absolute path for %s: %w", path, err)
		}
		cfg.WatchPaths[i] = absPath
	}

	return &cfg, nil
}

package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	WatchPaths []string `mapstructure:"watch"`
	Ignore     []string `mapstructure:"ignore"`
	DBPath     string   `mapstructure:"db_path"`
	MaxSize    int64    `mapstructure:"max_size"`
	Workers    int      `mapstructure:"workers"`
}

func Load() (*Config, error) {
	viper.SetConfigName("hippo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(os.Getenv("HOME") + "/.hippo")

	viper.SetDefault("watch", []string{"."})
	viper.SetDefault("ignore", []string{".git", "node_modules", "dist", "vendor"})
	viper.SetDefault("db_path", "hippo.db")
	viper.SetDefault("max_size", 10*1024*1024) // 10MB default
	viper.SetDefault("workers", runtime.NumCPU())

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &cfg, nil
}

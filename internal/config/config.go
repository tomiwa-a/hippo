package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	WatchPaths []string `mapstructure:"watch"`
	Ignore     []string `mapstructure:"ignore"`
	DBPath     string   `mapstructure:"db_path"`
}

func Load() (*Config, error) {
	viper.SetConfigName("hippo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(os.Getenv("HOME") + "/.hippo")

	viper.SetDefault("watch", []string{"."})
	viper.SetDefault("ignore", []string{".git", "node_modules", "dist", "vendor"})
	viper.SetDefault("db_path", "hippo.db")

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

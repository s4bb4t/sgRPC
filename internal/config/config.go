package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env      string
	Port     string
	TokenTTL time.Duration
	Timeout  time.Duration
}

func MustLoad() *Config {
	configPath := os.Getenv("CFG_PATH")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("config file not found: %s", configPath)
	}

	return &cfg
}

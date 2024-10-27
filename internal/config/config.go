package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	TokenTTL    time.Duration `yaml:"token_ttl" env-default:"1h"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	Env         string        `yaml:"env" env-default:"local"`
	Port        int
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

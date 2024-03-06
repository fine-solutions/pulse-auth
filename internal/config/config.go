package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"dev"`
	HttpServer `yaml:"http_server"`
	DB         Storage
}

type HttpServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Storage struct {
	Name     string `env:"DB_NAME"`
	Port     int    `env:"DB_PORT"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	configPath := getEnv("CONFIG_PATH", "")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("configs file doesn't exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can't read configs from %s: %v", configPath, err)
	}
	return &cfg

}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

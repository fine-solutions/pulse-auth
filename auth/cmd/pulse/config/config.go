package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"pulse-auth/internal/utils"
	"time"
)

type Config struct {
	Logger       LoggerConfig      `yaml:"logger"`
	Application  ApplicationConfig `yaml:"application"`
	PublicServer ServerConfig      `yaml:"public_server"`
	AdminServer  ServerConfig      `yaml:"admin_server"`
	Storage      StorageConfig
}

type LoggerConfig struct {
	Level       utils.Level       `yaml:"level"`
	Encoding    string            `yaml:"encoding"`
	Path        string            `yaml:"path"`
	Environment utils.Environment `yaml:"environment"`
}

type ApplicationConfig struct {
	GracefulShutdownTimeout time.Duration `yaml:"graceful_shutdown_timeout"`
	App                     string        `yaml:"app"`
	SaltValue               string        `yaml:"salt_value"`
}

type ServerConfig struct {
	Enable       bool   `yaml:"enable"`
	Endpoint     string `yaml:"endpoint"`
	Port         int    `yaml:"port" env:"PORT"`
	JwtTokenSalt string `yaml:"jwt_token_salt" env:"JWT_TOKEN_SALT"`
}

type StorageConfig struct {
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

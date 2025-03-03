package config

import (
	"os"
	"time"
)

type Config struct {
	MongoDB MongoDBConfig
}

type MongoDBConfig struct {
	URI            string
	Database       string
	ConnectTimeout time.Duration
	MaxPoolSize    uint64
	MinPoolSize    uint64
}

type HTTPConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{
		MongoDB: MongoDBConfig{
			URI:            getEnv("MONGODB_URI", "mongodb://localhost:27017"),
			Database:       getEnv("MONGODB_DATABASE", "goath"),
			ConnectTimeout: time.Second * 10,
			MaxPoolSize:    100,
			MinPoolSize:    10,
		},
		HTTP: HTTPConfig{
			Port:         getEnv("HTTP_PORT", "3111"),
			ReadTimeout:  time.Second * 15,
			WriteTimeout: time.Second * 15,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

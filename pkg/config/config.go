package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Debug               bool   `envconfig:"DEBUG" default:"false"`
	FirebaseCredentials string `envconfig:"FIREBASE_ACCOUNT_KEY" required:"true"`
	PostgresHost        string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresPort        int    `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresUser        string `envconfig:"POSTGRES_USER" default:"upmeet"`
	PostgresPassword    string `envconfig:"POSTGRES_PASSWORD" default:"upmeet"`
	PostgresDatabase    string `envconfig:"POSTGRES_DATABASE" default:"upmeet"`
	PostgresSSLMode     string `envconfig:"POSTGRES_SSLMODE" default:"disable"`
	BindAddress         string `envconfig:"BIND_ADDRESS" default:":3000"`
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	var cfg Config
	err := envconfig.Process("UPMEET", &cfg)
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	return &cfg
}

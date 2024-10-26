package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Database `mapstructure:",squash"`
	Logger   `mapstructure:",squash"`
}

func NewConfig(configFile string) *Config {
	var config Config
	if err := godotenv.Load(configFile); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatalf("Error reading environment variables: %v", err)
	}

	return &config
}

type Database struct {
	DBHost   string `env:"DBHOST"`
	DbUser   string `env:"DBUSER"`
	DbPass   string `env:"DBPASS"`
	DbPort   string `env:"DBPORT"`
	DbName   string `env:"DBNAME"`
	DbSchema string `env:"DBSCHEMA"`

	Debug bool `env:"DEBUG" env-default:"false"`
}

type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

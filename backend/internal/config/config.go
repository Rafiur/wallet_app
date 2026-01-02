package config

import (
	//"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

//type Config struct {
//	Database `mapstructure:",squash"`
//	Logger   `mapstructure:",squash"`
//}
//
//func NewConfig(configFile string) *Config {
//	var config Config
//	if err := godotenv.Load(configFile); err != nil {
//		log.Printf("Error loading .env file: %v", err)
//	}
//	if err := cleanenv.ReadEnv(&config); err != nil {
//		log.Fatalf("Error reading environment variables: %v", err)
//	}
//
//	return &config
//}

//type Database struct {
//	DBHost   string `env:"DBHOST"`
//	DbUser   string `env:"DBUSER"`
//	DbPass   string `env:"DBPASS"`
//	DbPort   string `env:"DBPORT"`
//	DbName   string `env:"DBNAME"`
//	DbSchema string `env:"DBSCHEMA"`
//
//	Debug bool `env:"DEBUG" env-default:"false"`
//}
//
//type Logger struct {
//	Development       bool
//	DisableCaller     bool
//	DisableStacktrace bool
//	Encoding          string
//	Level             string
//}

type Config struct {
	ServerPort    string
	JwtSecret     string
	JwtAccessTTL  time.Duration
	JwtRefreshTTL time.Duration

	DBHost   string
	DBUser   string
	DBPass   string
	DBPort   string
	DBName   string
	DBSchema string
	Debug    bool

	LogDevelopment       bool
	LogDisableCaller     bool
	LogDisableStacktrace bool
	LogEncoding          string
	LogLevel             string
}

var dynamicConfig *Config

func Init() {
	if err := godotenv.Load("config.env"); err != nil {
		log.Println("No config.env file found, relying on system env vars")
	}

	jwtAccessTTL, _ := time.ParseDuration(getEnv("JWT_ACCESS_TTL", "15m"))
	jwtRefreshTTL, _ := time.ParseDuration(getEnv("JWT_REFRESH_TTL", "168h"))

	debug, _ := strconv.ParseBool(getEnv("DEBUG", "false"))

	dynamicConfig = &Config{
		ServerPort:    getEnv("HTTP_ADDRESS", "8000"),
		JwtSecret:     getEnv("JWT_SECRET", "super-secret-jwt-key-change-me"), // ADD THIS TO config.env !!
		JwtAccessTTL:  jwtAccessTTL,
		JwtRefreshTTL: jwtRefreshTTL,

		DBHost:   getEnv("DBHOST", "localhost"),
		DBUser:   getEnv("DBUSER", "postgres"),
		DBPass:   getEnv("DBPASS", "admin1234"),
		DBPort:   getEnv("DBPORT", "5432"),
		DBName:   getEnv("DBNAME", "wallet_db"),
		DBSchema: getEnv("DBSCHEMA", "public"),
		Debug:    debug,

		LogDevelopment:       getEnv("LOG_DEVELOPMENT", "true") == "true",
		LogDisableCaller:     getEnv("LOG_DISABLE_CALLER", "false") == "true",
		LogDisableStacktrace: getEnv("LOG_DISABLE_STACKTRACE", "false") == "true",
		LogEncoding:          getEnv("LOG_ENCODING", "json"),
		LogLevel:             getEnv("LOG_LEVEL", "debug"),
	}
}

func GetDynamicConfig() *Config {
	return dynamicConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

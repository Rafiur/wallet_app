package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ServerPort    string
	JwtSecret     string
	JwtAccessTTL  time.Duration
	JwtRefreshTTL time.Duration

	CorsOrigins     []string
	RateLimitWindow time.Duration
	RateLimitMax    int

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
	rateLimitWindow, _ := time.ParseDuration(getEnv("RATE_LIMIT_WINDOW", "1m"))
	rateLimitMax, _ := strconv.Atoi(getEnv("RATE_LIMIT_MAX", "120"))

	corsOrigins := strings.Split(getEnv("CORS_ORIGINS", "http://localhost:5173,http://localhost:3000"), ",")
	for i := range corsOrigins {
		corsOrigins[i] = strings.TrimSpace(corsOrigins[i])
	}

	debug, _ := strconv.ParseBool(getEnv("DEBUG", "false"))

	dynamicConfig = &Config{
		ServerPort:    getEnv("HTTP_ADDRESS", "8000"),
		JwtSecret:     getEnv("JWT_SECRET", "super-secret-jwt-key-change-me"), // ADD THIS TO config.env !!
		JwtAccessTTL:  jwtAccessTTL,
		JwtRefreshTTL: jwtRefreshTTL,

		CorsOrigins:     corsOrigins,
		RateLimitWindow: rateLimitWindow,
		RateLimitMax:    rateLimitMax,

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

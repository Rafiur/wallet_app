package main

import (
	"fmt"
	"github.com/Rafiur/wallet_app/internal/config"
	"github.com/Rafiur/wallet_app/internal/config/database/postgres"
	"github.com/Rafiur/wallet_app/pkg/logger"
)

func main() {
	cfg := config.NewConfig("config.env")
	postgresDB := postgres.NewPostgresDB(cfg)
	log := logger.NewApiLogger(cfg)
	log.InitLogger()
	fmt.Println(postgresDB)

}

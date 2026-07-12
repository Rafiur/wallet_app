package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Rafiur/wallet_app/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"
	"log"
)

var DB *bun.DB

func NewPostgresDB() (*bun.DB, error) {
	conf := config.GetDynamicConfig()
	if conf == nil {
		return nil, fmt.Errorf("config not initialized")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName, conf.DBSchema)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	if err := sqldb.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	if conf.Debug {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	db.AddQueryHook(bunotel.NewQueryHook())

	DB = db
	log.Println("Connected to Postgres successfully")
	return db, nil
}

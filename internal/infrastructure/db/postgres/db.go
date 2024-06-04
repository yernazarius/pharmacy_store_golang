package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yernazarius/pharmacy_store_golang/configs"
)

func NewPostgresDB(config *configs.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

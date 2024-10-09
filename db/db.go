package db

import (
	"database/sql"
	logger "ssproxy/back/internal/pkg"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		logger.Error.Println(err)
	}

	return db, nil
}

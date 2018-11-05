package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func SetDSN(dsn string) error {
	var err error
	DB, err = sqlx.Open("mysql", dsn)

	if err != nil {
		return err
	}

	return DB.Ping()
}

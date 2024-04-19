package data

import (
	"database/sql"
	"sn-api/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); erro != nil {
		db.Close()
	}
}

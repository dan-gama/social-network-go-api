package data

import (
	"database/sql"
	"fmt"
	"sn-api/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)
	fmt.Println(config.ConnectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

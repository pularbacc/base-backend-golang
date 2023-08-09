package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"pular.server/env"
)

func Connect() (*sql.DB, error) {
	var db *sql.DB

	// Get a database handle.
	var err error
	db, err = sql.Open("postgres", env.DB_CONNECT_STR)
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	fmt.Println("DB Connected!")

	return db, nil
}

package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func New(databaseDriver, databaseConnectionString string) (*sql.DB, error) {
	var err error
	db, err := sql.Open(databaseDriver, databaseConnectionString)
	if err != nil {
		return &sql.DB{}, err
	}
	return db, nil
}

package infrastructure

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func NewDB(dbName string) (db *sql.DB,err error) {
	db, err = sql.Open("sqlite", fmt.Sprintf("%s?cache=shared", dbName))
	if err != nil {
		return nil, err
	}
	return db, nil
}

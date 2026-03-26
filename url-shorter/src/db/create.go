package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func initTable(db *sql.DB) error {
	query:=`CREATE TABLE IF NOT EXISTS urls (
  id TEXT PRIMARY KEY,
  url TEXT NOT NULL,
  short_url TEXT NOT NULL,
  count INTEGER DEFAULT 0
);`
	_, err := db.Exec(query)
	return err
}
var DB *sql.DB
func ConnectDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite", "test.db")
    if err != nil {
        return nil, err
    }
    if err := initTable(db); err != nil {
        return nil, err
    }
    DB = db
    return db, nil
}


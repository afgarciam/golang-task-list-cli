package ado

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Initialize() {
	db, err := GetConnection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt := `CREATE TABLE IF NOT EXISTS todo (
		id INTEGER PRIMARY KEY,
		task TEXT NOT NULL,
		complete INTEGER DEFAULT 0,
		created_at INTEGER DEFAULT (strftime('%s','now'))
	)`

	_, err = db.Exec(stmt)

	if err != nil {
		log.Fatal(err)
	}
}

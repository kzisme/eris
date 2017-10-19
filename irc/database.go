package irc

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// OpenDB opens a connection to a sqlite3 database given a path
func OpenDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("open db error: ", err)
	}
	return db
}

package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// File to use for the database
const file string = "timeclock.db"

// Create the database if needed
// TODO: register an admin user
const create string = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY,
    name STRING,
    email STRING,
    password STRING,
    admin BOOLEAN
);

CREATE TABLE IF NOT EXISTS entries (
    id INTEGER NOT NULL PRIMARY KEY,
    FOREIGN KEY(userid) REFERENCES users(id) ON DELETE CASCADE,
    start DATETIME,
    end DATETIME,
    note STRING
);`

func New() (*sql.DB, error) {
	// Add `?mode=memory` for in-memory version, `?mode=rwc` for the real file
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return db, nil
}

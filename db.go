package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// File to use for the database
const file string = "file:timeclock.db"

// Create the database if needed
// TODO: register an admin user here?
const create string = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY,
    name STRING,
    email STRING,
    password STRING,
    salt STRING,
    admin BOOLEAN
);

CREATE TABLE IF NOT EXISTS entries (
    id INTEGER NOT NULL PRIMARY KEY,
    start DATETIME,
    end DATETIME,
    note STRING,
    userid INTEGER,
    FOREIGN KEY(userid) REFERENCES users(id)
);`

const DB_MODE_MEM string = "?mode=memory"
const DB_MODE_FILE string = "?mode=rwc"

func NewDB(mode string) (*sql.DB, error) {
	if mode != DB_MODE_MEM && mode != DB_MODE_FILE {
		return nil, fmt.Errorf("mode %s not supported", mode)
	}
	// Add `?mode=memory` for in-memory version, `?mode=rwc` for the real file
	db, err := sql.Open("sqlite3", file+mode)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return db, nil
}

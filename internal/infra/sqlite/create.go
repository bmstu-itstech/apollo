package sqlite

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"os"
)

// MustLoad loads or (if non-existent) creates an SQLiteStorage file
func MustLoad(dataFile string) SQLiteStorage {
	if _, err := os.Stat(dataFile); errors.Is(err, os.ErrNotExist) {
		os.Create(dataFile)
		slog.Debug("created sqlite3 database file", "file", dataFile)
	}
	db, err := sql.Open("sqlite3", dataFile)
	if err != nil {
		slog.Error("error opening sqlite3 database", "error", err)
		os.Exit(1)
	}

	// created_at is a unix timestamp
	const create_materials string = `
	CREATE TABLE "materials" (
		"uuid"	TEXT NOT NULL UNIQUE,
		"name"	TEXT NOT NULL,
		"description"	TEXT,
		"url"	TEXT NOT NULL,
		"author"	TEXT,
		"views" INTEGER NOT NULL,
		"department_id" INTEGER NOT NULL,
		"discipline_id" INTEGER NOT NULL,
		"created_at" INTEGER NOT NULL,
		PRIMARY KEY("uuid")
	);`
	statement, err := db.Prepare(create_materials)
	if err != nil {
		if err.Error() != "table \"materials\" already exists" {
			slog.Error("error initialising sqlite3 database table", "error", err)
			os.Exit(1)
		}
	} else {
		statement.Exec()
	}

	const create_departments string = `
	CREATE TABLE "departments" (
		"id" INTEGER NOT NULL UNIQUE,
		"name"	TEXT NOT NULL,
		"description"	TEXT NOT NULL,
		PRIMARY KEY("id")
	);`
	statement, err = db.Prepare(create_departments)
	if err != nil {
		if err.Error() != "table \"departments\" already exists" {
			slog.Error("error initialising sqlite3 database table", "error", err)
			os.Exit(1)
		}
	} else {
		statement.Exec()
	}

	const create_disciplines string = `
	CREATE TABLE "disciplines" (
		"id" INTEGER NOT NULL UNIQUE,
		"name"	TEXT NOT NULL,
		PRIMARY KEY("id")
	);`
	statement, err = db.Prepare(create_disciplines)
	if err != nil {
		if err.Error() != "table \"disciplines\" already exists" {
			slog.Error("error initialising sqlite3 database table", "error", err)
			os.Exit(1)
		}
	} else {
		statement.Exec()
	}

	slog.Debug("loaded database", "file", dataFile)
	return SQLiteStorage{db}
}

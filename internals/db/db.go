package db

import (
	"database/sql"
	"gonote/internals/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const schema string = "CREATE TABLE IF NOT EXISTS todos (id TEXT PRIMARY KEY, content TEXT, completed INTEGER)"

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", config.DatabaseURL)

	if err != nil {
		log.Fatal("Unable to start Database", err)
		return nil
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Unable to ping Database", err)
		return nil
	}

	if _, err = db.Exec(schema); err != nil {
		log.Fatal("Unable to create Database", err)
		return nil
	}

	return db
}

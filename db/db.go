package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var appdb *sql.DB

func init() {
	s, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error connecting to the database: " + err.Error())
	}
	appdb = s
}

func AppDB() *sql.DB {
	return appdb
}

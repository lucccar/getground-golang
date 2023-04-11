package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbUser   = "gosql"
	dbPass   = "root"
	dbName   = "goserver"
)

type DBConnector interface {
	Close() error
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sql.Row
	Query(query string, args ...any) (*sql.Rows, error)
}

func Connect() DBConnector {

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	PORT := os.Getenv("DB_PORT")

	connStr := USER + ":" + PASS + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME
	// USER+":"+PASS+"@/"+DBNAME

	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	return db
}

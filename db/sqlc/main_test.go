package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
	dbsource = "postgresql://root:password@localhost:5432/simplebank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		log.Fatal("cannot connect to database", err)

	}
	testQueries = New(conn)

	os.Exit(m.Run())
}

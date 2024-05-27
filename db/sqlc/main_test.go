package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Cannot load variables ", err)
	}

	dbdriver := os.Getenv("dbdriver")
	dbsource := os.Getenv("dbsource")
	var err error
	testDB, err = sql.Open(dbdriver, dbsource)
	if err != nil {
		log.Fatal("cannot connect to database", err)

	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}

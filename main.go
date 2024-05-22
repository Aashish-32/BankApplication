package main

import (
	"database/sql"
	"log"

	api "github.com/Aashish-32/bank/api"
	db "github.com/Aashish-32/bank/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {

	const (
		dbdriver = "postgres"
		dbsource = "postgresql://root:password@localhost:5432/simplebank?sslmode=disable"
	)
	DB, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		log.Fatal("cannot connect to database", err)

	}
	store := db.NewStore(DB)
	server := api.NewServer(store)
	err = server.Start("localhost:8000")
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}

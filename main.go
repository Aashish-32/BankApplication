package main

import (
	"database/sql"
	"log"
	"os"

	api "github.com/Aashish-32/bank/api"
	db "github.com/Aashish-32/bank/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("cannot load enviroment variables ", err)
	}

	dbdriver := os.Getenv("dbdriver")
	dbsource := os.Getenv("dbsource")
	server_address := os.Getenv("server_address")

	DB, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		log.Fatal("cannot connect to database ", err)

	}
	store := db.NewStore(DB)
	server := api.NewServer(store)

	err = server.Start(server_address)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}

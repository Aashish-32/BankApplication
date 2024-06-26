package main

import (
	"database/sql"
	"log"
	"os"

	api "github.com/Aashish-32/bank/api"
	db "github.com/Aashish-32/bank/db/sqlc"
	"github.com/Aashish-32/bank/util"
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
	token_key := os.Getenv("token_symmetric_key")

	DB, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		log.Fatal("cannot connect to database ", err)

	}
	store := db.NewStore(DB)
	var config = util.Config{
		TokenSymmetricKey: token_key,
	}
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalln(err)
	}

	err = server.Start(server_address)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}

package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/simplebank/api"
	db "github.com/simplebank/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	var err error

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}

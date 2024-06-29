package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/musstafayilmazz/Golang/simplebank/api"
	db "github.com/musstafayilmazz/Golang/simplebank/db/sqlc"

	"log"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAdress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Can not connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}

}

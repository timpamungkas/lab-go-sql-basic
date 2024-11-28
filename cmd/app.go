package main

import (
	"database/sql"
	"log"

	"github.com/alpharent/apartment/datastore"
)

func main() {
	datastore.InitializeDatabase("db/migration", "sqlite3://alpharent.db")

	db, err := sql.Open("sqlite", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	clientCount, _ := datastore.CountClients(db)
	log.Printf("Number of clients: %d\n", clientCount)

	apartmentCount, _ := datastore.CountClientApartments(db)
	log.Printf("Number of apartments: %d\n", apartmentCount)
}

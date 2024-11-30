package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/alpharent/apartment/datastore"
	filemigration "github.com/alpharent/apartment/file_migration"
	httphandler "github.com/alpharent/apartment/http_handler"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	datastore.InitializeDatabase("db/migration", "sqlite3://alpharent.db")

	var err error
	db, err = sql.Open("sqlite", "alpharent.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	clientStore := datastore.NewClientStore(db)
	clientApartmentStore := datastore.NewClientApartmentStore(db)
	clientHttpHandler := httphandler.NewClientHttpHandler(*clientStore)
	clientApartmentHttpHandler := httphandler.NewClientApartmentHttpHandler(*clientApartmentStore)

	migrationFromFile := filemigration.NewMigrationFromFile(*clientStore, *clientApartmentStore)
	migrationFromFile.MigrateClientsFromCsv("db/csv/clients.csv", true)
	migrationFromFile.MigrateClientApartmentsFromCsv("db/csv/client_apartments.csv", true)
	migrationFromFile.MigrateClientsFromJson("db/json/clients.json")
	migrationFromFile.MigrateClientApartmentsFromJson("db/json/client_apartments.json")

	http.HandleFunc("/api/count/clients", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		clientHttpHandler.CountClientsHandler(w, r)
	})

	http.HandleFunc("/api/count/client-apartments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		clientApartmentHttpHandler.CountClientApartmentsHandler(w, r)
	})

	http.HandleFunc("/api/transaction/client", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		clientHttpHandler.InsertClientHandler(w, r)
	})

	http.HandleFunc("/api/clients/all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		clientHttpHandler.SelectAllClientsHandler(w, r)
	})

	http.HandleFunc("/api/transaction/client-apartment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		clientApartmentHttpHandler.InsertClientApartmentHandler(w, r)
	})

	http.HandleFunc("/api/client-apartments/all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		clientApartmentHttpHandler.SelectAllClientApartmentsHandler(w, r)
	})

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

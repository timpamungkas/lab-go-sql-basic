package main

import (
	"github.com/alpharent/apartment/datastore"
)

func main() {
	datastore.InitializeDatabase("db/migration", "sqlite3://alpharent.db")
}

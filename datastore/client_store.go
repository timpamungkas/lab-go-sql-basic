package datastore

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

const SQL_COUNT_CLIENTS = `SELECT COUNT(client_id) FROM clients;`

func CountClients(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow(SQL_COUNT_CLIENTS).Scan(&count)

	if err != nil {
		return -1, err
	}

	return count, nil
}

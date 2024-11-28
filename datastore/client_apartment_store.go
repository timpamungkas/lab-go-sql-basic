package datastore

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

const SQL_COUNT_CLIENT_APARTMENTS = `SELECT COUNT(apartment_id) FROM client_apartments;`

func CountClientApartments(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow(SQL_COUNT_CLIENT_APARTMENTS).Scan(&count)

	if err != nil {
		return -1, err
	}

	return count, nil
}

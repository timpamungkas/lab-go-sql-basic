package datastore

import (
	"database/sql"
)

type ClientStore struct {
	db *sql.DB
}

func NewClientStore(db *sql.DB) *ClientStore {
	return &ClientStore{db: db}
}

const sqlCountClients = `SELECT COUNT(*) FROM clients;`

func (s *ClientStore) CountClients() (int, error) {
	var count int
	err := s.db.QueryRow(sqlCountClients).Scan(&count)

	if err != nil {
		return -1, err
	}

	return count, nil
}

type ClientDatabaseRow struct {
	ClientID string         `json:"client_id"`
	FullName string         `json:"full_name"`
	Email    sql.NullString `json:"email"`
	Phone    string         `json:"phone"`
}

const sqlInsertClient = `
	INSERT INTO clients(
		client_id, full_name, email, phone
	)
	VALUES(
		$1, $2, $3, $4
	)
`

func (s *ClientStore) InsertClient(client *ClientDatabaseRow) error {
	_, err := s.db.Exec(
		sqlInsertClient,
		client.ClientID, client.FullName, client.Email, client.Phone)

	return err
}

const sqlSelectAllClients = `
	SELECT client_id, full_name, email, phone
	FROM clients;
`

func (s *ClientStore) SelectAllClients() ([]ClientDatabaseRow, error) {
	rows, err := s.db.Query(sqlSelectAllClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := []ClientDatabaseRow{}
	for rows.Next() {
		var client ClientDatabaseRow
		err := rows.Scan(&client.ClientID, &client.FullName, &client.Email, &client.Phone)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

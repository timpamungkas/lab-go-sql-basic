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
	ClientID string
	FullName string
	Email    sql.NullString
	Phone    string
}

const sqlInsertClient = `
	INSERT INTO clients(
		client_id, full_name, email, phone
	)
	VALUES(
		$1, $2, NULLIF($3,''), $4
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
	    FROM clients
	ORDER BY full_name;
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

type ClientWithApartmentsDatabaseRow struct {
	ClientID   string
	FullName   string
	Email      sql.NullString
	Phone      string
	Apartments []ClientApartmentDatabaseRow
}

const sqlSelectAllClientsWithApartments = `
	   SELECT c.client_id, c.full_name, c.email, c.phone, 
	  		  a.apartment_id, a.description, a.building_name, a.room_number, a.street_address,
	          a.city, a.postal_code, a.is_available_for_rent, a.rent_price
	     FROM clients c
	LEFT JOIN apartments a 
			  ON c.client_id = a.client_id
`

func (s *ClientStore) SelectAllClientsWithApartments() (
	map[ClientWithApartmentsDatabaseRow][]ClientApartmentDatabaseRow, error) {
	rows, err := s.db.Query(sqlSelectAllClientsWithApartments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clientsMap := make(map[string]*ClientDatabaseRow)
	for rows.Next() {
		var clientID, fullName, email, phone string
		var apartmentID, description, buildingName, roomNumber, streetAddress, city, postalCode sql.NullString
		var isAvailableForRent sql.NullBool
		var rentPrice sql.NullFloat64

		err := rows.Scan(&clientID, &fullName, &email, &phone,
			&apartmentID, &description, &buildingName, &roomNumber,
			&streetAddress, &city, &postalCode, &isAvailableForRent, &rentPrice)
		if err != nil {
			return nil, err
		}

		client, exists := clientsMap[clientID]
		if !exists {
			client = &ClientDatabaseRow{
				ClientID: clientID,
				FullName: fullName,
				Phone:    phone,
			}

			if email != "" {
				client.Email = sql.NullString{String: email, Valid: true}
			}

			clientsMap[clientID] = client
		}

		apartment := ClientApartmentDatabaseRow{
			ApartmentID:        apartmentID.String,
			Description:        description.String,
			BuildingName:       buildingName.String,
			RoomNumber:         roomNumber.String,
			StreetAddress:      streetAddress.String,
			City:               city.String,
			PostalCode:         postalCode.String,
			IsAvailableForRent: isAvailableForRent.Bool,
			RentPrice:          rentPrice.Float64,
		}
		client.Apartments = append(client.Apartments, apartment)
	}

	clients := make([]Client, 0, len(clientsMap))
	for _, client := range clientsMap {
		clients = append(clients, *client)
	}

	return clients, nil
}

package datastore

import (
	"database/sql"
)

type ClientApartmentStore struct {
	db *sql.DB
}

func NewClientApartmentStore(db *sql.DB) *ClientApartmentStore {
	return &ClientApartmentStore{db: db}
}

const sqlCountClientApartments = `SELECT COUNT(apartment_id) FROM client_apartments;`

func (s *ClientApartmentStore) CountClientApartments() (int, error) {
	var count int
	err := s.db.QueryRow(sqlCountClientApartments).Scan(&count)

	if err != nil {
		return -1, err
	}

	return count, nil
}

type ClientApartmentDatabaseRow struct {
	ApartmentID        string
	Description        sql.NullString
	BuildingName       sql.NullString
	RoomNumber         sql.NullString
	StreetAddress      string
	City               string
	PostalCode         sql.NullString
	IsAvailableForRent sql.NullBool
	RentPrice          float64
	ClientID           string
}

const sqlInsertClientApartment = `
	INSERT INTO client_apartments(
		apartment_id, description, building_name, room_number, street_address,
		city, postal_code, is_available_for_rent, rent_price, client_id
	)
	VALUES(
		$1, NULLIF($2, ''), 
		NULLIF($3, ''), NULLIF($4, ''), 
		$5, $6, 
		NULLIF($7, ''), 
		$8, $9, 
		$10
	)
`

func (s *ClientApartmentStore) InsertClientApartment(clientApartment *ClientApartmentDatabaseRow) error {
	_, err := s.db.Exec(
		sqlInsertClientApartment,
		clientApartment.ApartmentID, clientApartment.Description,
		clientApartment.BuildingName, clientApartment.RoomNumber,
		clientApartment.StreetAddress, clientApartment.City,
		clientApartment.PostalCode, clientApartment.IsAvailableForRent,
		clientApartment.RentPrice, clientApartment.ClientID)

	return err
}

const sqlSelectAllClientApartments = `
	  SELECT apartment_id, description, building_name, room_number, street_address,
	       	 city, postal_code, is_available_for_rent, rent_price, client_id
	    FROM client_apartments
	ORDER BY city;
`

func (s *ClientApartmentStore) SelectAllClientApartments() ([]ClientApartmentDatabaseRow, error) {
	rows, err := s.db.Query(sqlSelectAllClientApartments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clientApartments := []ClientApartmentDatabaseRow{}
	for rows.Next() {
		var clientApartment ClientApartmentDatabaseRow
		err := rows.Scan(&clientApartment.ApartmentID, &clientApartment.Description,
			&clientApartment.BuildingName, &clientApartment.RoomNumber,
			&clientApartment.StreetAddress, &clientApartment.City,
			&clientApartment.PostalCode, &clientApartment.IsAvailableForRent,
			&clientApartment.RentPrice, &clientApartment.ClientID)
		if err != nil {
			return nil, err
		}
		clientApartments = append(clientApartments, clientApartment)
	}

	return clientApartments, nil
}

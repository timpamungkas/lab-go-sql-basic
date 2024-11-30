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

// type ClientApartmentHttp struct {
// 	ApartmentID        *string  `json:"apartment_id"`
// 	Description        *string  `json:"description"`
// 	BuildingName       *string  `json:"building_name"`
// 	RoomNumber         *string  `json:"room_number"`
// 	StreetAddress      *string  `json:"street_address"`
// 	City               *string  `json:"city"`
// 	PostalCode         *string  `json:"postal_code"`
// 	IsAvailableForRent *bool    `json:"is_available_for_rent"`
// 	RentPrice          *float64 `json:"rent_price"`
// 	ClientID           *string  `json:"client_id"`
// }

// const sqlInsertClientApartment = `
// 	INSERT INTO client_apartments(
// 		apartment_id, description, building_name, room_number, street_address,
// 		city, postal_code, is_available_for_rent, rent_price, client_id
// 	)
// 	VALUES(
// 		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
// 	)
// `

// func InsertClientApartment(db *sql.DB, clientApartment *ClientApartmentHttp) error {
// 	_, err := db.Exec(
// 		sqlInsertClientApartment,
// 		clientApartment.ApartmentID, clientApartment.Description,
// 		clientApartment.BuildingName, clientApartment.RoomNumber,
// 		clientApartment.StreetAddress, clientApartment.City,
// 		clientApartment.PostalCode, clientApartment.IsAvailableForRent,
// 		clientApartment.RentPrice, clientApartment.ClientID)

// 	return err
// }

// type ClientApartmentDatabaseRow struct {
// 	ApartmentID        string         `json:"apartment_id"`
// 	Description        sql.NullString `json:"description"`
// 	BuildingName       sql.NullString `json:"building_name"`
// 	RoomNumber         sql.NullString `json:"room_number"`
// 	StreetAddress      string         `json:"street_address"`
// 	City               string         `json:"city"`
// 	PostalCode         sql.NullString `json:"postal_code"`
// 	IsAvailableForRent sql.NullBool   `json:"is_available_for_rent"`
// 	RentPrice          float64        `json:"rent_price"`
// 	ClientID           string         `json:"client_id"`
// }

// const sqlSelectAllClientApartments = `
// 	SELECT apartment_id, description, building_name, room_number, street_address,
// 		city, postal_code, is_available_for_rent, rent_price, client_id
// 	FROM client_apartments;
// `

// func SelectAllClientApartments(db *sql.DB) ([]ClientApartmentDatabaseRow, error) {
// 	rows, err := db.Query(sqlSelectAllClientApartments)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	clientApartments := []ClientApartmentDatabaseRow{}
// 	for rows.Next() {
// 		var clientApartment ClientApartmentDatabaseRow
// 		err := rows.Scan(&clientApartment.ApartmentID, &clientApartment.Description,
// 			&clientApartment.BuildingName, &clientApartment.RoomNumber,
// 			&clientApartment.StreetAddress, &clientApartment.City,
// 			&clientApartment.PostalCode, &clientApartment.IsAvailableForRent,
// 			&clientApartment.RentPrice, &clientApartment.ClientID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		clientApartments = append(clientApartments, clientApartment)
// 	}

// 	return clientApartments, nil
// }

package filemigration

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alpharent/apartment/datastore"
	httphandler "github.com/alpharent/apartment/http_handler"
)

type MigrationFromFile struct {
	clientStore          datastore.ClientStore
	clientApartmentStore datastore.ClientApartmentStore
}

func NewMigrationFromFile(clientStore datastore.ClientStore,
	clientApartmentStore datastore.ClientApartmentStore) *MigrationFromFile {
	return &MigrationFromFile{
		clientStore:          clientStore,
		clientApartmentStore: clientApartmentStore,
	}
}

func (m *MigrationFromFile) MigrateClientsFromCsv(migrationFile string, skipHeader bool) {
	csvFile, err := os.Open(migrationFile)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	if skipHeader && len(records) > 1 {
		records = records[1:]
	}

	for _, record := range records {
		clientDatabaseRow := datastore.ClientDatabaseRow{
			ClientID: record[0],
			FullName: record[1],
			Email:    sql.NullString{String: record[2], Valid: true},
			Phone:    record[3],
		}

		m.clientStore.InsertClient(&clientDatabaseRow)
	}
}

func parseFloat(value string) float64 {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return floatValue
}

func (m *MigrationFromFile) MigrateClientApartmentsFromCsv(
	migrationFile string, skipHeader bool) {
	csvFile, err := os.Open(migrationFile)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	if skipHeader && len(records) > 1 {
		records = records[1:]
	}

	for _, record := range records {
		clientApartmentDatabaseRow := datastore.ClientApartmentDatabaseRow{
			ApartmentID:        record[0],
			Description:        sql.NullString{String: record[1], Valid: true},
			BuildingName:       sql.NullString{String: record[2], Valid: true},
			RoomNumber:         sql.NullString{String: record[3], Valid: true},
			StreetAddress:      record[4],
			City:               record[5],
			PostalCode:         sql.NullString{String: record[6], Valid: true},
			IsAvailableForRent: sql.NullBool{Bool: strings.ToLower(record[7]) == "true", Valid: true},
			RentPrice:          parseFloat(record[8]),
			ClientID:           record[9],
		}

		m.clientApartmentStore.InsertClientApartment(&clientApartmentDatabaseRow)
	}
}

func (h *MigrationFromFile) convertToClientDatabaseRow(
	clientHttp httphandler.ClientHttp) datastore.ClientDatabaseRow {
	clientDatabaseRow := datastore.ClientDatabaseRow{
		ClientID: *clientHttp.ClientID,
		FullName: *clientHttp.FullName,
		Phone:    *clientHttp.Phone,
	}

	if clientHttp.Email != nil {
		clientDatabaseRow.Email = sql.NullString{String: *clientHttp.Email, Valid: true}
	}

	return clientDatabaseRow
}

func (m *MigrationFromFile) MigrateClientsFromJson(migrationFile string) {
	jsonFile, err := os.Open(migrationFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	var records []httphandler.ClientHttp
	err = decoder.Decode(&records)
	if err != nil {
		fmt.Println(err)
	}

	for _, record := range records {
		databaseRow := m.convertToClientDatabaseRow(record)
		m.clientStore.InsertClient(&databaseRow)
	}
}

func (m *MigrationFromFile) convertToClientApartmentDatabaseRow(
	clientApartmentHttp httphandler.ClientApartmentHttp) datastore.ClientApartmentDatabaseRow {
	clientApartmentDatabaseRow := datastore.ClientApartmentDatabaseRow{
		ApartmentID:   *clientApartmentHttp.ApartmentID,
		StreetAddress: *clientApartmentHttp.StreetAddress,
		City:          *clientApartmentHttp.City,
		RentPrice:     *clientApartmentHttp.RentPrice,
		ClientID:      *clientApartmentHttp.ClientID,
	}

	if clientApartmentHttp.Description != nil {
		clientApartmentDatabaseRow.Description = sql.NullString{String: *clientApartmentHttp.Description, Valid: true}
	}

	if clientApartmentHttp.BuildingName != nil {
		clientApartmentDatabaseRow.BuildingName = sql.NullString{String: *clientApartmentHttp.BuildingName, Valid: true}
	}

	if clientApartmentHttp.RoomNumber != nil {
		clientApartmentDatabaseRow.RoomNumber = sql.NullString{String: *clientApartmentHttp.RoomNumber, Valid: true}
	}

	if clientApartmentHttp.PostalCode != nil {
		clientApartmentDatabaseRow.PostalCode = sql.NullString{String: *clientApartmentHttp.PostalCode, Valid: true}
	}

	if clientApartmentHttp.IsAvailableForRent != nil {
		clientApartmentDatabaseRow.IsAvailableForRent = sql.NullBool{Bool: *clientApartmentHttp.IsAvailableForRent, Valid: true}
	}

	return clientApartmentDatabaseRow
}

func (m *MigrationFromFile) MigrateClientApartmentsFromJson(migrationFile string) {
	jsonFile, err := os.Open(migrationFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	var records []httphandler.ClientApartmentHttp
	err = decoder.Decode(&records)
	if err != nil {
		fmt.Println(err)
	}

	for _, record := range records {
		databaseRow := m.convertToClientApartmentDatabaseRow(record)
		m.clientApartmentStore.InsertClientApartment(&databaseRow)
	}
}

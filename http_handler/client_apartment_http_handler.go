package httphandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alpharent/apartment/datastore"
)

type ClientApartmentHttpHandler struct {
	datastore datastore.ClientApartmentStore
}

func NewClientApartmentHttpHandler(datastore datastore.ClientApartmentStore) *ClientApartmentHttpHandler {
	return &ClientApartmentHttpHandler{datastore: datastore}
}

func (h *ClientApartmentHttpHandler) CountClientApartmentsHandler(w http.ResponseWriter, r *http.Request) {
	count, err := h.datastore.CountClientApartments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]int{"client_apartments_count": count}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

type ClientApartmentHttp struct {
	ApartmentID        *string  `json:"apartment_id"`
	Description        *string  `json:"description"`
	BuildingName       *string  `json:"building_name"`
	RoomNumber         *string  `json:"room_number"`
	StreetAddress      *string  `json:"street_address"`
	City               *string  `json:"city"`
	PostalCode         *string  `json:"postal_code"`
	IsAvailableForRent *bool    `json:"is_available_for_rent"`
	RentPrice          *float64 `json:"rent_price"`
	ClientID           *string  `json:"client_id"`
}

func (h *ClientApartmentHttpHandler) convertToClientApartmentDatabaseRow(clientApartmentHttp ClientApartmentHttp) datastore.ClientApartmentDatabaseRow {
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

func (h *ClientApartmentHttpHandler) convertToClientApartmentHttp(clientApartmentRow datastore.ClientApartmentDatabaseRow) ClientApartmentHttp {
	clientApartmentHttp := ClientApartmentHttp{
		ApartmentID:        &clientApartmentRow.ApartmentID,
		StreetAddress:      &clientApartmentRow.StreetAddress,
		City:               &clientApartmentRow.City,
		RentPrice:          &clientApartmentRow.RentPrice,
		ClientID:           &clientApartmentRow.ClientID,
		IsAvailableForRent: nil,
	}

	if clientApartmentRow.Description.Valid {
		clientApartmentHttp.Description = &clientApartmentRow.Description.String
	}

	if clientApartmentRow.BuildingName.Valid {
		clientApartmentHttp.BuildingName = &clientApartmentRow.BuildingName.String
	}

	if clientApartmentRow.RoomNumber.Valid {
		clientApartmentHttp.RoomNumber = &clientApartmentRow.RoomNumber.String
	}

	if clientApartmentRow.PostalCode.Valid {
		clientApartmentHttp.PostalCode = &clientApartmentRow.PostalCode.String
	}

	if clientApartmentRow.IsAvailableForRent.Valid {
		clientApartmentHttp.IsAvailableForRent = &clientApartmentRow.IsAvailableForRent.Bool
	}

	return clientApartmentHttp
}

func (h *ClientApartmentHttpHandler) InsertClientApartmentHandler(w http.ResponseWriter, r *http.Request) {
	var clientApartmentHttp ClientApartmentHttp

	if err := json.NewDecoder(r.Body).Decode(&clientApartmentHttp); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	clientApartmentDatabaseRow := h.convertToClientApartmentDatabaseRow(clientApartmentHttp)

	if err := h.datastore.InsertClientApartment(&clientApartmentDatabaseRow); err != nil {
		errMessage := fmt.Sprintf("Failed to insert client apartment: %v", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var responseMessage = fmt.Sprintf(
		"Client apartment %s has been successfully inserted",
		*clientApartmentHttp.ApartmentID)
	response := map[string]string{"message": responseMessage}

	json.NewEncoder(w).Encode(response)
}

func (h *ClientApartmentHttpHandler) SelectAllClientApartmentsHandler(w http.ResponseWriter, r *http.Request) {
	clientApartments, err := h.datastore.SelectAllClientApartments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientApartmentsHttp := []ClientApartmentHttp{}

	for _, clientApartment := range clientApartments {
		clientApartmentHttp := h.convertToClientApartmentHttp(clientApartment)
		clientApartmentsHttp = append(clientApartmentsHttp, clientApartmentHttp)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(clientApartmentsHttp)
}

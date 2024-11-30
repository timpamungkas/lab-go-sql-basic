package httphandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alpharent/apartment/datastore"
)

type ClientHttpHandler struct {
	datastore datastore.ClientStore
}

func NewClientHttpHandler(datastore datastore.ClientStore) *ClientHttpHandler {
	return &ClientHttpHandler{datastore: datastore}
}

func (h *ClientHttpHandler) CountClientsHandler(w http.ResponseWriter, r *http.Request) {
	count, err := h.datastore.CountClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]int{"client_count": count}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

type ClientHttp struct {
	ClientID *string `json:"client_id"`
	FullName *string `json:"full_name"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

func (h *ClientHttpHandler) convertToClientDatabaseRow(clientHttp ClientHttp) datastore.ClientDatabaseRow {
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

func (h *ClientHttpHandler) convertToClientHttp(clientRow datastore.ClientDatabaseRow) ClientHttp {
	clientHttp := ClientHttp{
		ClientID: &clientRow.ClientID,
		FullName: &clientRow.FullName,
		Phone:    &clientRow.Phone,
	}

	if clientRow.Email.Valid {
		clientHttp.Email = &clientRow.Email.String
	}

	return clientHttp
}

func (h *ClientHttpHandler) InsertClientHandler(w http.ResponseWriter, r *http.Request) {
	var clientHttp ClientHttp

	if err := json.NewDecoder(r.Body).Decode(&clientHttp); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	clientDatabaseRow := h.convertToClientDatabaseRow(clientHttp)

	if err := h.datastore.InsertClient(&clientDatabaseRow); err != nil {
		errMessage := fmt.Sprintf("Failed to insert client: %v", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var responseMessage = fmt.Sprintf("Client %s has been successfully inserted", *clientHttp.ClientID)
	response := map[string]string{"message": responseMessage}

	json.NewEncoder(w).Encode(response)
}

func (h *ClientHttpHandler) SelectAllClientsHandler(w http.ResponseWriter, r *http.Request) {
	clients, err := h.datastore.SelectAllClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientsHttp := []ClientHttp{}

	for _, clientRow := range clients {
		clientHttp := h.convertToClientHttp(clientRow)
		clientsHttp = append(clientsHttp, clientHttp)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(clientsHttp)
}

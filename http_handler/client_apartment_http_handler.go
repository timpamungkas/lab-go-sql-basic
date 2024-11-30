package httphandler

import (
	"encoding/json"
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

// func insertClientApartmentHandler(w http.ResponseWriter, r *http.Request) {
// 	var clientApartment datastore.ClientApartmentHttp

// 	if err := json.NewDecoder(r.Body).Decode(&clientApartment); err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	if err := datastore.InsertClientApartment(db, &clientApartment); err != nil {
// 		errMessage := fmt.Sprintf("Failed to insert client apartment: %v", err)
// 		http.Error(w, errMessage, http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)

// 	var responseMessage = fmt.Sprintf("Client apartment %s has been successfully inserted", *clientApartment.ApartmentID)
// 	response := map[string]string{"message": responseMessage}

// 	json.NewEncoder(w).Encode(response)
// }

// func selectAllClientApartmentsHandler(w http.ResponseWriter, r *http.Request) {
// 	clientApartments, err := datastore.SelectAllClientApartments(db)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(clientApartments)
// }

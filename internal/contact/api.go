package contact

import (
	"net/http"

	"github.com/gorilla/mux"
)

type resource struct {
}

func RegisterHandler(r *mux.Router) {
	resource := resource{}
	r.HandleFunc("/contacts", resource.query).Methods(http.MethodGet)
	// r.HandleFunc("/contact", createContact).Methods(http.MethodPost)
}

func (resource resource) query(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"name":"ram", "phone":"1232424"}`))
}

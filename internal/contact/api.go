package contact

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sub-rat/contactbook-api/pkg/utils"
)

type resource struct {
	service ServiceInterface
}

func RegisterHandler(r *mux.Router, service ServiceInterface) {
	resource := resource{service: service}
	r.HandleFunc("/contacts", resource.query).Methods(http.MethodGet)
	// r.HandleFunc("/contact", createContact).Methods(http.MethodPost)
}

func (resource resource) query(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 20
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	offset := (page - 1) * limit
	contactList, err := resource.service.Query(offset, limit, query)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, contactList)
}

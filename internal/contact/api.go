package contact

import (
	"encoding/json"
	"io/ioutil"
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
	r.HandleFunc("/contact", resource.create).Methods(http.MethodPost)
	r.HandleFunc("/contact/{id}", resource.get).Methods(http.MethodGet)
}

func (resource resource) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	contact, err := resource.service.Get(uint(id))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, contact)
}

func (resource resource) create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	data := Contact{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		utils.RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	contact, err := resource.service.Create(data)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, contact)
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

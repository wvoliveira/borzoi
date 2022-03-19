package user

import (
	"encoding/json"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/gorilla/mux"
	"net/http"
)

func (s service) HTTPNew(r *mux.Router) {
	rr := r.PathPrefix("/v1/users").Subrouter()

	//r.OPTIONS("", nil)
	rr.HandleFunc("/{id}", s.HTTPFindByID).Methods("GET")
	rr.HandleFunc("/{id}", s.HTTPUpdate).Methods("PATCH")
}

func (s service) HTTPFindByID(w http.ResponseWriter, r *http.Request) {
	req, err := decodeFindByID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	user, err := s.FindByID(req.ID)
	if err != nil {
		handleError(w, err)
		return
	}

	err = encodeFindByID(w, user)
	if err != nil {
		handleError(w, err)
		return
	}
	return
}

func (s service) HTTPUpdate(w http.ResponseWriter, r *http.Request) {
	req, err := decodeUpdate(r)
	if err != nil {
		handleError(w, err)
		return
	}

	l, err := s.Update(req.ID, entity.User{Name: req.Name})
	if err != nil {
		handleError(w, err)
		return
	}

	err = encodeUpdate(w, l)
	if err != nil {
		handleError(w, err)
		return
	}
	return
}

func handleError(w http.ResponseWriter, err error) {
	r := response{Status: "error", Data: nil, Message: err.Error()}
	w.Header().Set("Content-Type", "application/json")

	// TODO: change this for switch errors.
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(r)
}

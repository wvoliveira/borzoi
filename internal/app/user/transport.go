package user

import (
	"encoding/json"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	m "github.com/elga-io/borzoi/internal/pkg/middleware"
	res "github.com/elga-io/borzoi/internal/pkg/response"
	e "github.com/elga-io/canideos/errors"
	"github.com/gorilla/mux"
	"net/http"
)

func (s service) HTTPNew(r *mux.Router) {
	rr := r.PathPrefix("/v1/users").Subrouter()

	rr.Use(m.Middleware{Cache: s.cache}.Auth)

	rr.HandleFunc("", s.HTTPFindAll).Methods("GET")
	rr.HandleFunc("/{id}", s.HTTPFindByID).Methods("GET")
	rr.HandleFunc("/{id}", s.HTTPUpdate).Methods("PATCH")
	rr.HandleFunc("/me", s.HTTPFindMe).Methods("GET")
	rr.HandleFunc("/me", s.HTTPUpdateMe).Methods("PATCH")
}

func (s service) HTTPFindAll(w http.ResponseWriter, r *http.Request) {
	res.NotImplemented(w)
	return
}

func (s service) HTTPFindByID(w http.ResponseWriter, r *http.Request) {
	req, err := decodeFindByID(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	user, err := s.FindByID(req.ID)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeFindByID(w, user)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

func (s service) HTTPUpdate(w http.ResponseWriter, r *http.Request) {
	req, err := decodeUpdate(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	l, err := s.Update(req.ID, entity.User{Name: req.Name})
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeUpdate(w, l)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

func (s service) HTTPFindMe(w http.ResponseWriter, r *http.Request) {
	res.NotImplemented(w)
	return
}

func (s service) HTTPUpdateMe(w http.ResponseWriter, r *http.Request) {
	res.NotImplemented(w)
	return
}

func handleError(w http.ResponseWriter, err error) {
	r := res.Response{Status: "error", Data: nil, Message: err.Error()}
	w.Header().Set("Content-Type", "application/json")

	// TODO: change this for switch errors.
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(r)
}

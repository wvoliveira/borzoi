package user

import (
	"github.com/elga-io/borzoi/internal/pkg/entity"
	e "github.com/elga-io/borzoi/internal/pkg/errors"
	m "github.com/elga-io/borzoi/internal/pkg/middleware"
	res "github.com/elga-io/borzoi/internal/pkg/response"
	"github.com/gorilla/mux"
	"net/http"
)

func (s service) HTTPNew(r *mux.Router) {
	rr := r.PathPrefix("/v1/users").Subrouter()
	rr.Use(m.Middleware{Cache: s.cache}.Auth)

	rr.HandleFunc("", s.HTTPFindAll).Methods("GET")
	rr.HandleFunc("/me", s.HTTPFindMe).Methods("GET")
	rr.HandleFunc("/me", s.HTTPUpdateMe).Methods("PATCH")
	rr.HandleFunc("/{id}", s.HTTPFindByID).Methods("GET")
	rr.HandleFunc("/{id}", s.HTTPUpdate).Methods("PATCH")
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

	user, err := s.FindByID(r.Context(), req.ID)
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

	l, err := s.Update(r.Context(), req.ID, entity.User{Name: req.Name})
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
	userID, err := decodeFindMe(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	user, err := s.FindMe(r.Context(), userID)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeFindMe(w, user)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

func (s service) HTTPUpdateMe(w http.ResponseWriter, r *http.Request) {
	res.NotImplemented(w)
	return
}

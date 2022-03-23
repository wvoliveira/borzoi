package client

import (
	"github.com/elga-io/borzoi/internal/pkg/entity"
	m "github.com/elga-io/borzoi/internal/pkg/middleware"
	e "github.com/elga-io/canideos/errors"
	"github.com/gorilla/mux"
	"net/http"
)

func (s service) HTTPNew(r *mux.Router) {
	rr := r.PathPrefix("/v1/clients").Subrouter()
	rr.Use(m.Middleware{Cache: s.cache}.Auth)

	rr.HandleFunc("", s.HTTPFindAll).Methods("GET")
	rr.HandleFunc("", s.HTTPCreate).Methods("POST")
	rr.HandleFunc("/{id}", s.HTTPFindByID).Methods("GET")
	rr.HandleFunc("/{id}", s.HTTPUpdate).Methods("PATCH")
	rr.HandleFunc("/{id}", s.HTTPDelete).Methods("DELETE")
}

func (s service) HTTPFindAll(w http.ResponseWriter, r *http.Request) {
	search, page, limit, err := decodeFindAll(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	clients, err := s.FindAll(r.Context(), search, page, limit)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeFindAll(w, clients)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

func (s service) HTTPCreate(w http.ResponseWriter, r *http.Request) {
	client, err := decodeCreate(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = s.Create(r.Context(), client)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeCreate(w)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

func (s service) HTTPFindByID(w http.ResponseWriter, r *http.Request) {
	id, err := decodeFindByID(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	user, err := s.FindByID(r.Context(), id)
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

	l, err := s.Update(r.Context(), req.ID, entity.Client{Name: req.Name})
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

func (s service) HTTPDelete(w http.ResponseWriter, r *http.Request) {
	id, del, err := decodeDelete(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	client, err := s.Delete(r.Context(), id, del)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeDelete(w, client)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

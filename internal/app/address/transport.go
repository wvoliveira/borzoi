package address

import (
	e "github.com/elga-io/borzoi/internal/pkg/errors"
	m "github.com/elga-io/borzoi/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func (s service) HTTPNew(r *mux.Router) {
	rr := r.PathPrefix("/v1/addresses").Subrouter()
	rr.Use(m.Middleware{Cache: s.cache}.Auth)

	rr.HandleFunc("", s.HTTPFindAll).Methods("GET")
	rr.HandleFunc("", s.HTTPCreate).Methods("POST")
	rr.HandleFunc("/{id}", s.HTTPFindByID).Methods("GET")
	rr.HandleFunc("/{id}", s.HTTPUpdate).Methods("PATCH")
	rr.HandleFunc("/{id}", s.HTTPDelete).Methods("DELETE")
}

func (s service) HTTPFindAll(w http.ResponseWriter, r *http.Request) {
	search, clientID, page, limit, err := decodeFindAll(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	addresses, err := s.FindAll(r.Context(), search, clientID, page, limit)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeFindAll(w, addresses)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

func (s service) HTTPCreate(w http.ResponseWriter, r *http.Request) {
	address, err := decodeCreate(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = s.Create(r.Context(), address)
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

	address, err := s.FindByID(r.Context(), id)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeFindByID(w, address)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

func (s service) HTTPUpdate(w http.ResponseWriter, r *http.Request) {
	id, action, address, clients, err := decodeUpdate(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = s.Update(r.Context(), id, action, address, clients)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeUpdate(w)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

func (s service) HTTPDelete(w http.ResponseWriter, r *http.Request) {
	id, err := decodeDelete(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	address, err := s.Delete(r.Context(), id)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	err = encodeDelete(w, address)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	return
}

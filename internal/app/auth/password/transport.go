package password

import (
	"github.com/elga-io/borzoi/internal/pkg/entity"
	e "github.com/elga-io/borzoi/internal/pkg/errors"
	"github.com/gorilla/mux"
	"net/http"
)

func (s service) HTTPNew(r *mux.Router) {
	// TODO: middleware to check cookie authentication.
	rr := r.PathPrefix("/auth/password").Subrouter()

	rr.HandleFunc("/login", s.HTTPLogin).Methods("POST")
	rr.HandleFunc("/register", s.HTTPRegister).Methods("POST")
}

func (s service) HTTPLogin(w http.ResponseWriter, r *http.Request) {
	// Decode request to request object.
	dr, err := decodeLoginRequest(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	identity := entity.Identity{Provider: "email", UID: dr.Email, Password: dr.Password}

	// Business logic.
	token, err := s.Login(r.Context(), identity)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	// Encode response to send to client.
	err = encodeLogin(w, token)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
}

func (s service) HTTPRegister(w http.ResponseWriter, r *http.Request) {
	// Decode request to request object.
	dr, err := decodeRegisterRequest(r)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
	identity := entity.Identity{Provider: "email", UID: dr.Email, Password: dr.Password}
	user := entity.User{Name: dr.Name}

	// Business logic.
	err = s.Register(r.Context(), identity, user)
	if err != nil {
		e.EncodeError(w, err)
		return
	}

	// Encode object to answer request (response).
	err = encodeRegister(w)
	if err != nil {
		e.EncodeError(w, err)
		return
	}
}

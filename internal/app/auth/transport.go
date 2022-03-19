package auth

import (
	"fmt"
	e "github.com/elga-io/canideos/errors"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func (s service) HTTPNew(r *mux.Router) {
	// TODO add middlewares to check authentication and authorization.
	rr := r.PathPrefix("/auth").Subrouter()

	rr.HandleFunc("/logout", s.HTTPLogout).Methods("GET")
}

func (s service) HTTPLogout(w http.ResponseWriter, r *http.Request) {
	// Decode request to object.
	token, err := decodeLogout(r)
	if err != nil {
		fmt.Printf("%s error: decode logout: %s\n", time.Now(), err.Error())
		e.EncodeError(w, err)
		return
	}

	// Business logic.
	err = s.Logout(nil, token)
	if err != nil {
		fmt.Printf("%s error: service logout: %s\n", time.Now(), err.Error())
		e.EncodeError(w, err)
		return
	}

	err = encodeLogout(w)
	if err != nil {
		fmt.Printf("%s error: service logout: %s\n", time.Now(), err.Error())
		e.EncodeError(w, err)
		return
	}
}

package auth

import (
	"fmt"
	m "github.com/elga-io/borzoi/internal/pkg/middleware"
	e "github.com/elga-io/canideos/errors"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (s service) HTTPNew(r *mux.Router) {
	rr := r.PathPrefix("/auth").Subrouter()
	rr.Use(m.Middleware{Cache: s.cache}.Auth)
	rr.HandleFunc("/logout", s.HTTPLogout).Methods("POST")
}

func (s service) HTTPLogout(w http.ResponseWriter, r *http.Request) {
	l := log.Ctx(r.Context())

	// Decode request to object.
	token, err := decodeLogout(r)
	if err != nil {
		l.Error().Caller().Msg(fmt.Sprintf("decode logout: %s", err.Error()))
		e.EncodeError(w, err)
		return
	}

	// Business logic.
	err = s.Logout(r.Context(), token)
	if err != nil {
		l.Error().Caller().Msg(fmt.Sprintf("service logout: %s", err.Error()))
		e.EncodeError(w, err)
		return
	}

	err = encodeLogout(w)
	if err != nil {
		l.Error().Caller().Msg(fmt.Sprintf("encode logout: %s", err.Error()))
		e.EncodeError(w, err)
		return
	}
}

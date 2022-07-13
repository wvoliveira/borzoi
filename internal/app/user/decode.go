package user

import (
	"encoding/json"
	"errors"
	e "github.com/elga-io/borzoi/internal/pkg/errors"
	"github.com/elga-io/borzoi/internal/pkg/session"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

type findByIDRequest struct {
	ID string `json:"id"`
}

type updateRequest struct {
	ID   string `json:"-"`
	Name string `json:"name"`
}

// GET /v1/users/{id}
func decodeFindByID(r *http.Request) (req findByIDRequest, err error) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		return req, errors.New("you need pass user_id. Ex.: api/v1/user/123e4567-e89b-12d3-a456-426655440000")
	}
	return
}

// PATCH /v1/users/{id}
func decodeUpdate(r *http.Request) (req updateRequest, err error) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		return req, errors.New("you need pass user_id. Ex.: api/v1/user/123e4567-e89b-12d3-a456-426655440000")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if err.Error() == "EOF" {
			return req, errors.New("you need send a body/payload for update user information")
		}
		if err != nil {
			return
		}
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return
	}
	return
}

// GET /v1/users/me
func decodeFindMe(r *http.Request) (userID string, err error) {
	l := log.Ctx(r.Context())
	userID = session.UserGetIDFromContext(r.Context())

	if userID == "" {
		l.Warn().Caller().Msg("user ID not found from session cookie")
		return userID, e.ErrUserNotFound
	}
	return
}

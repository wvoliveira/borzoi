package client

import (
	"encoding/json"
	"errors"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	e "github.com/elga-io/canideos/errors"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

// GET /v1/clients
func decodeFindAll(r *http.Request) (search string, page, limit int, err error) {
	return
}

// PUT /v1/clients
func decodeCreate(r *http.Request) (client entity.Client, err error) {
	err = json.NewDecoder(r.Body).Decode(&client)
	if err == io.EOF {
		return client, e.ErrClientBadRequest
	}
	return client, err
}

// GET /v1/clients/{id}
func decodeFindByID(r *http.Request) (id string, err error) {
	vars := mux.Vars(r)
	id = vars["id"]
	if id == "" {
		return id, errors.New("you need pass a id. Ex.: api/v1/clients/123e4567-e89b-12d3-a456-426655440000")
	}
	return
}

// PATCH /v1/clients/{id}
func decodeUpdate(r *http.Request) (client entity.Client, err error) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		return client, errors.New("you need pass a id. Ex.: api/v1/clients/123e4567-e89b-12d3-a456-426655440000")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if err.Error() == "EOF" {
			return client, errors.New("you need send a body for update client information")
		}
		if err != nil {
			return
		}
	}
	err = json.Unmarshal(body, &client)
	if err != nil {
		return
	}
	return
}

// DELETE /v1/clients/{id}
func decodeDelete(r *http.Request) (id string, err error) {
	return
}

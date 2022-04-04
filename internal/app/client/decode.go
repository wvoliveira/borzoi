package client

import (
	"encoding/json"
	"errors"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	e "github.com/elga-io/borzoi/internal/pkg/errors"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type createRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GET /v1/clients
func decodeFindAll(r *http.Request) (search string, page, limit int, err error) {
	params := r.URL.Query()
	paramSearch := params.Get("q")
	search = strings.TrimSpace(paramSearch)

	paramPage := params.Get("page")
	if paramPage == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			page = 1
		}
		if page < 1 {
			page = 1
		}
	}

	paramLimit := params.Get("limit")
	if paramLimit == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(paramLimit)
		if err != nil {
			limit = 10
		}
	}
	return
}

// POST /v1/clients
func decodeCreate(r *http.Request) (client entity.Client, err error) {
	cr := createRequest{}
	err = json.NewDecoder(r.Body).Decode(&cr)
	if err == io.EOF {
		return client, e.ErrClientBadRequest
	}
	if cr.Name == "" || cr.Description == "" {
		return client, e.ErrClientBadRequest
	}
	client.Name = cr.Name
	client.Description = cr.Description
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
	vars := mux.Vars(r)
	id = vars["id"]
	if id == "" {
		return id, errors.New("you need pass a id. Ex.: api/v1/clients/123e4567-e89b-12d3-a456-426655440000")
	}
	return id, nil
}

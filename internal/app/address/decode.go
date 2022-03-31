package address

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

type CreateRequest struct {
	Country    string `json:"country"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	CEP        string `json:"cep"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`
}

// GET /v1/addresses
func decodeFindAll(r *http.Request) (search, clientID string, page, limit int, err error) {
	params := r.URL.Query()
	search = strings.TrimSpace(params.Get("q"))
	clientID = strings.TrimSpace(params.Get("client_id"))

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

// PUT /v1/addresses
func decodeCreate(r *http.Request) (a entity.Address, err error) {
	cq := CreateRequest{}
	err = json.NewDecoder(r.Body).Decode(&cq)
	if err == io.EOF {
		return a, e.ErrClientBadRequest
	}
	if err != nil {
		return
	}

	a.Country = cq.Country
	a.Name = cq.Name
	a.Phone = cq.Phone
	a.Street = cq.Street
	a.Complement = cq.Complement
	a.District = cq.District
	a.City = cq.City
	a.State = cq.State

	if v, err := strconv.Atoi(cq.CEP); err == nil {
		a.CEP = v
	}
	if v, err := strconv.Atoi(cq.Number); err == nil {
		a.Number = v
	}
	return
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

// PATCH /v1/addresses/{id}
func decodeUpdate(r *http.Request) (address entity.Address, err error) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		return address, errors.New("you must pass a id. Ex.: api/v1/addresses/123e4567-e89b-12d3-a456-426655440000")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if err.Error() == "EOF" {
			return address, errors.New("you must send a body for update the address")
		}
		if err != nil {
			return
		}
	}
	err = json.Unmarshal(body, &address)
	if err != nil {
		return
	}
	return
}

// DELETE /v1/addresses/{id}
func decodeDelete(r *http.Request) (id, clientID string, err error) {
	vars := mux.Vars(r)
	id = vars["id"]
	if id == "" {
		return id, clientID, errors.New("you must pass a id. Ex.: api/v1/addresses/123e4567-e89b-12d3-a456-426655440000")
	}

	params := r.URL.Query()
	clientID = strings.TrimSpace(params.Get("client_id"))
	return
}

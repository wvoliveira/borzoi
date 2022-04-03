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

type UpdateRequest struct {
	Action    string        `json:"action"`
	Address   CreateRequest `json:"address"`
	ClientIDs []string      `json:"client_ids"`
}

func (c CreateRequest) ToAddress() (addr entity.Address) {
	c.Country = addr.Country
	c.Name = addr.Name
	c.Phone = addr.Phone
	c.Street = addr.Street
	c.Complement = addr.Complement
	c.District = addr.District
	c.City = addr.City
	c.State = addr.State
	if v, err := strconv.Atoi(c.CEP); err == nil {
		addr.CEP = v
	}
	if v, err := strconv.Atoi(c.Number); err == nil {
		addr.Number = v
	}
	return
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

// POST /v1/addresses
func decodeCreate(r *http.Request) (addr entity.Address, err error) {
	cq := CreateRequest{}
	err = json.NewDecoder(r.Body).Decode(&cq)
	if err == io.EOF {
		return addr, e.ErrClientBadRequest
	}
	if err != nil {
		return
	}
	addr = cq.ToAddress()
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
// Actions:
// - empty: update only address
// - append: append clients from address
// - remove: remove clients from address
func decodeUpdate(r *http.Request) (id, action string, addr entity.Address, clientIDs []string, err error) {
	vars := mux.Vars(r)
	id = vars["id"]

	if id == "" {
		return id, action, addr, clientIDs, errors.New("you must pass a id. Ex.: api/v1/addresses/<id>")
	}

	req := UpdateRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if err.Error() == "EOF" {
			return id, action, addr, clientIDs, errors.New("you must send a body for update the address")
		}
		if err != nil {
			return
		}
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		return
	}

	addr = req.Address.ToAddress()
	return id, req.Action, addr, req.ClientIDs, nil
}

// DELETE /v1/addresses/{id}
func decodeDelete(r *http.Request) (id string, err error) {
	vars := mux.Vars(r)
	id = vars["id"]
	if id == "" {
		return id, errors.New("you must pass a id. Ex.: api/v1/addresses/<id>")
	}
	return
}

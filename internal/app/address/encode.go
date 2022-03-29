package client

import (
	"fmt"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/borzoi/internal/pkg/response"
	"net/http"
)

// GET /v1/clients
func encodeFindAll(w http.ResponseWriter, clients []entity.Client) (err error) {
	fmt.Println("clients")
	fmt.Println(clients)
	response.Normal(w, clients, "", http.StatusOK)
	return
}

// PUT /v1/clients
func encodeCreate(w http.ResponseWriter) (err error) {
	response.Normal(w, nil, "", http.StatusCreated)
	return
}

// GET /v1/clients/{id}
func encodeFindByID(w http.ResponseWriter, client entity.Client) (err error) {
	response.Normal(w, client, "", http.StatusOK)
	return
}

// PATCH /v1/clients/{id}
func encodeUpdate(w http.ResponseWriter, client entity.Client) (err error) {
	response.Normal(w, client, "", http.StatusOK)
	return
}

// DELETE /v1/clients/{id}
func encodeDelete(w http.ResponseWriter, client entity.Client) (err error) {
	response.Normal(w, client, "", http.StatusOK)
	return
}

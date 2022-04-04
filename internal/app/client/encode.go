package client

import (
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/borzoi/internal/pkg/response"
	"net/http"
)

// GET /v1/clients
func encodeFindAll(w http.ResponseWriter, clients []entity.Client, page, pages, limit int, total int64) (err error) {
	type findAll struct {
		Page  int         `json:"page"`
		Pages int         `json:"pages"`
		Limit int         `json:"per_page"`
		Total int64       `json:"total"`
		Data  interface{} `json:"clients"`
	}
	data := findAll{page, pages, limit, total, clients}
	response.Default(w, data, "", http.StatusOK)
	return
}

// PUT /v1/clients
func encodeCreate(w http.ResponseWriter) (err error) {
	response.Default(w, nil, "", http.StatusCreated)
	return
}

// GET /v1/clients/{id}
func encodeFindByID(w http.ResponseWriter, client entity.Client) (err error) {
	response.Default(w, client, "", http.StatusOK)
	return
}

// PATCH /v1/clients/{id}
func encodeUpdate(w http.ResponseWriter, client entity.Client) (err error) {
	response.Default(w, client, "", http.StatusOK)
	return
}

// DELETE /v1/clients/{id}
func encodeDelete(w http.ResponseWriter) (err error) {
	response.Default(w, nil, "", http.StatusNoContent)
	return
}

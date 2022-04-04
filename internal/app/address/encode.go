package address

import (
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/borzoi/internal/pkg/response"
	"net/http"
)

// GET /v1/addresses
func encodeFindAll(w http.ResponseWriter, addresses []entity.Address) (err error) {
	response.Default(w, addresses, "", http.StatusOK)
	return
}

// PUT /v1/addresses
func encodeCreate(w http.ResponseWriter) (err error) {
	response.Default(w, nil, "", http.StatusCreated)
	return
}

// GET /v1/addresses/{id}
func encodeFindByID(w http.ResponseWriter, address entity.Address) (err error) {
	response.Default(w, address, "", http.StatusOK)
	return
}

// PATCH /v1/addresses/{id}
func encodeUpdate(w http.ResponseWriter) (err error) {
	response.Default(w, nil, "", http.StatusOK)
	return
}

// DELETE /v1/addresses/{id}
func encodeDelete(w http.ResponseWriter, address entity.Address) (err error) {
	response.Default(w, address, "", http.StatusOK)
	return
}

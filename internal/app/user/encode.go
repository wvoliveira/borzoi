package user

import (
	"net/http"

	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/borzoi/internal/pkg/response"
)

// GET /v1/user/<id>
func encodeFindByID(w http.ResponseWriter, user entity.User) {
	response.Default(w, user, "", http.StatusOK)
}

// PATCH /v1/users/<id>
func encodeUpdate(w http.ResponseWriter, user entity.User) {
	response.Default(w, nil, "", http.StatusOK)
}

// GET /v1/users/me
func encodeFindMe(w http.ResponseWriter, user entity.User) {
	response.Default(w, user, "", http.StatusOK)
}

// PATCH /v1/users/me
func encodeUpdateMe(w http.ResponseWriter) {
	response.Default(w, nil, "", http.StatusOK)
}

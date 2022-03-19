package user

import (
	"encoding/json"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"net/http"
)

type response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

func encodeFindByID(w http.ResponseWriter, user entity.User) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response{
		Status:  "successful",
		Data:    user,
		Message: "",
	})
	return
}

func encodeUpdate(w http.ResponseWriter, link entity.User) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response{
		Status:  "successful",
		Data:    link,
		Message: "",
	})
	return
}

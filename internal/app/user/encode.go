package user

import (
	"encoding/json"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	res "github.com/elga-io/borzoi/internal/pkg/response"
	"net/http"
)

func encodeFindByID(w http.ResponseWriter, user entity.User) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res.Response{
		Status:  "successful",
		Data:    user,
		Message: "",
	})
	return
}

func encodeUpdate(w http.ResponseWriter, link entity.User) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res.Response{
		Status:  "successful",
		Data:    link,
		Message: "",
	})
	return
}

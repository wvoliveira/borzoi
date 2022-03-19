package user

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type findByIDRequest struct {
	ID string `json:"id"`
}

type updateRequest struct {
	ID   string `json:"-"`
	Name string `json:"name"`
}

func decodeFindByID(r *http.Request) (req findByIDRequest, err error) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		return req, errors.New("you need pass user_id. Ex.: api/v1/user/123e4567-e89b-12d3-a456-426655440000")
	}
	return
}

func decodeUpdate(r *http.Request) (req updateRequest, err error) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		return req, errors.New("you need pass user_id. Ex.: api/v1/user/123e4567-e89b-12d3-a456-426655440000")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if err.Error() == "EOF" {
			return req, errors.New("you need send a body/payload for update user information")
		}
		if err != nil {
			return
		}
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return
	}
	return
}

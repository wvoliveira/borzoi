package password

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func decodeLoginRequest(r *http.Request) (req loginRequest, err error) {
	err = json.NewDecoder(r.Body).Decode(&req)
	if err == io.EOF {
		return req, errors.New("you need send a body with email and password fields")
	}
	return req, err
}

func decodeRegisterRequest(r *http.Request) (req registerRequest, err error) {
	err = json.NewDecoder(r.Body).Decode(&req)
	if err == io.EOF {
		return req, errors.New("you need send a body with name, email and password fields")
	}
	return req, err
}

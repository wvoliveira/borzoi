package auth

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func encodeLogout(w http.ResponseWriter) (err error) {
	cookie := &http.Cookie{
		Name:   "session",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response{
		Status: "successful",
	})
	return
}

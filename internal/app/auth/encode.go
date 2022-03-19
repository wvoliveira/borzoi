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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	cookie := &http.Cookie{
		Name:   "session",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	err = json.NewEncoder(w).Encode(response{
		Status: "successful",
	})
	return
}

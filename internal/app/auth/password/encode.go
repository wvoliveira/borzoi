package password

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func encodeLogin(w http.ResponseWriter, token string) (err error) {
	cookie := &http.Cookie{
		Name:  "session",
		Value: token,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response{
		Status: "successful",
	})
	return
}

func encodeRegister(w http.ResponseWriter) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response{
		Status: "successful",
	})
	return
}

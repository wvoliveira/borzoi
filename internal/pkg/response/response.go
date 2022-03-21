package response

import (
	"encoding/json"
	"net/http"
)

// Response default response for http requests.
type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

func NotImplemented(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Response{
		Status:  "successful",
		Message: "not implemented yet",
	})
}

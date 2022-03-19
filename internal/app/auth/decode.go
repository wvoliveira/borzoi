package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func decodeLogout(r *http.Request) (token string, err error) {
	c, err := r.Cookie("session")
	token = c.Value

	// TODO: remove it.
	fmt.Printf("%s debug: cookie session value: %s\n", time.Now(), token)

	if token == "" {
		return token, errors.New("token was not found")
	}
	return
}

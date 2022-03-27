package auth

import (
	"errors"
	"fmt"
	zl "github.com/rs/zerolog/log"
	"net/http"
)

func decodeLogout(r *http.Request) (token string, err error) {
	c, err := r.Cookie("session")
	token = c.Value
	zl.Debug().Caller().Msg(fmt.Sprintf("cookie session value: %s", token))
	if token == "" {
		return token, errors.New("token was not found")
	}
	return
}

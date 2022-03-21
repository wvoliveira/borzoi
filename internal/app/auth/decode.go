package auth

import (
	"errors"
	"fmt"
	zl "github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func decodeLogout(r *http.Request) (token string, err error) {
	c, err := r.Cookie("session")
	token = c.Value
	zl.Debug().Msg(fmt.Sprintf("%s debug: cookie session value: %s\n", time.Now(), token))
	if token == "" {
		return token, errors.New("token was not found")
	}
	return
}

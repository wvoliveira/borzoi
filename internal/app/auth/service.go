package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

// Service encapsulates the authentication logic.
type Service interface {
	Logout(ctx context.Context, token string) error

	HTTPNew(r *mux.Router)
	HTTPLogout(w http.ResponseWriter, r *http.Request)
}

type service struct {
	db    *gorm.DB
	cache *badger.DB
}

// NewService creates a new authentication service.
func NewService(db *gorm.DB, cache *badger.DB) Service {
	return service{db, cache}
}

// Logout remove cookie and refresh token from database.
func (s service) Logout(ctx context.Context, token string) (err error) {
	l := log.Ctx(ctx)

	key := fmt.Sprintf("auth/tokens/%s", token)
	err = s.cache.Update(func(txn *badger.Txn) (err error) {
		_, err = txn.Get([]byte(key))

		if err == badger.ErrKeyNotFound {
			l.Warn().Caller().Msg(fmt.Sprintf("token %s was not found", token))
			return errors.New("token was not found")
		}

		err = txn.Delete([]byte(key))
		return
	})

	if err != nil {
		l.Error().Caller().Msg(fmt.Sprintf("to delete token in cache: %s", err.Error()))
		return
	}
	return
}

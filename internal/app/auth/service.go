package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"time"
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
	fmt.Printf("%s info: starting logout for token %s\n", time.Now(), token)

	key := fmt.Sprintf("auth/tokens/%s", token)
	err = s.cache.Update(func(txn *badger.Txn) (err error) {
		_, err = txn.Get([]byte(key))

		if err == badger.ErrKeyNotFound {
			fmt.Printf("%s warn: token %s was not found", time.Now, token)
			return errors.New("token was not found")
		}

		err = txn.Delete([]byte(key))
		return
	})

	if err != nil {
		fmt.Printf("%s error: error to delete token in cache: %s", time.Now(), err.Error())
		return
	}
	return
}

package session

import (
	"context"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/rs/zerolog/log"
)

const (
	DBKey = "sessions/%s"
)

// UserGetIDFromContext get user struct from context.
func UserGetIDFromContext(ctx context.Context) (userID string) {
	if ctxUserID := ctx.Value("user_id"); ctxUserID != nil {
		return ctxUserID.(string)
	}
	return
}

// UserGetIDFromSession get id from noSQL database.
func UserGetIDFromSession(ctx context.Context, db *badger.DB, key []byte) (userID string, err error) {
	l := log.Ctx(ctx)

	err = db.View(func(txn *badger.Txn) (err error) {
		item, err := txn.Get(key)
		if err != nil {
			return
		}

		err = item.Value(func(val []byte) (err error) {
			userID = string(val)
			l.Debug().Caller().Msg(fmt.Sprintf("user_id from nosql: %s", userID))
			return
		})
		return
	})
	return
}

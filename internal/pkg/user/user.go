package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/rs/zerolog/log"
)

// FromContext get user struct from context.
func FromContext(ctx context.Context) entity.User {
	return ctx.Value("user").(entity.User)
}

// GetUserBySession get id from noSQL database.
func GetUserBySession(ctx context.Context, db *badger.DB, key []byte) (user entity.User, err error) {
	l := log.Ctx(ctx)

	err = db.View(func(txn *badger.Txn) (err error) {
		item, err := txn.Get(key)
		if err != nil {
			return
		}

		err = item.Value(func(val []byte) (err error) {
			err = json.Unmarshal(val, &user)
			if err != nil {
				l.Error().Caller().Msg(fmt.Sprintf("trying unmarshal user from badgerdb value - %s", err.Error()))
				return
			}
			l.Debug().Caller().Msg(fmt.Sprintf("user id from badgerdb: %s", user.ID))
			return
		})
		return
	})
	return
}

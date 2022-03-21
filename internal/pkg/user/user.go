package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	zl "github.com/rs/zerolog/log"
)

// FromContext get user struct from context.
func FromContext(ctx context.Context) entity.User {
	return ctx.Value("user").(entity.User)
}

// GetUserBySession get id from noSQL database.
func GetUserBySession(db *badger.DB, key []byte) (user entity.User, err error) {
	err = db.View(func(txn *badger.Txn) (err error) {
		item, err := txn.Get(key)
		if err != nil {
			return
		}
		err = item.Value(func(val []byte) (err error) {
			err = json.Unmarshal(val, &user)
			if err != nil {
				zl.Error().Caller().Msg(fmt.Sprintf("trying unmarshal user from badgerdb value - %s", err.Error()))
				return
			}
			zl.Debug().Caller().Msg(fmt.Sprintf("user id from badgerdb: %s", user.ID))
			return
		})
		return
	})
	return
}

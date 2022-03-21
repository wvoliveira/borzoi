package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"time"
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
				fmt.Printf("%s error: trying unmarshal user from badgerdb value - %s\n", time.Now(), err.Error())
				return
			}
			fmt.Printf("%s debug: user id from badgerdb: %s\n", time.Now(), user.ID)
			return
		})
		return
	})
	return
}

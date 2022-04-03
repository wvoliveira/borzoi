package address

import (
	"context"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	e "github.com/elga-io/borzoi/internal/pkg/errors"
	"github.com/elga-io/borzoi/internal/pkg/session"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

// Service encapsulates the link service logic, http handlers and another transport layer.
type Service interface {
	FindAll(ctx context.Context, search, clientID string, page, limit int) ([]entity.Address, error)
	FindByID(ctx context.Context, id string) (address entity.Address, err error)
	Update(ctx context.Context, id, action string, addr entity.Address, clientIDs []string) error
	Delete(ctx context.Context, id string) (address entity.Address, err error)

	HTTPNew(r *mux.Router)
	HTTPFindAll(w http.ResponseWriter, r *http.Request)
	HTTPFindByID(w http.ResponseWriter, r *http.Request)
	HTTPUpdate(w http.ResponseWriter, r *http.Request)
	HTTPDelete(w http.ResponseWriter, r *http.Request)
}

type service struct {
	db    *gorm.DB
	cache *badger.DB
}

// NewService creates a new authentication service.
func NewService(db *gorm.DB, cache *badger.DB) Service {
	return service{db, cache}
}

// FindAll get a list of clients.
func (s service) FindAll(ctx context.Context, search, clientID string, page, limit int) (addresses []entity.Address, err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)

	offset := 0
	if page > 1 {
		offset = page * limit
	}

	query := s.db.Model(&entity.Address{}).
		Debug().
		Limit(limit).
		Offset(offset)

	if search != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%[1]s%s%[1]s", "%", search))
	}

	if clientID != "" {
		err = query.
			Where("user_id = ? AND client.id = ?", userID, clientID).
			Preload("Clients").
			Find(&addresses).Error
	} else {
		err = query.
			Where("user_id = ?", userID).
			Preload("Clients").
			Find(&addresses).Error
	}

	if len(addresses) == 0 {
		l.Warn().Caller().Msg(fmt.Sprintf("clients with search=%s limit=%d offset=%d was not found", search, limit, offset))
		return addresses, nil
	}
	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	return
}

// Create add a new address.
func (s service) Create(ctx context.Context, address entity.Address) (err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)
	address.UserID = userID

	err = s.db.Debug().Model(&address).Create(&address).Save(&address).Error
	if err != nil {
		l.Error().Caller().Msg(err.Error())
	}
	return
}

// FindByID get a specific address by ID.
func (s service) FindByID(ctx context.Context, id string) (address entity.Address, err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)
	address.ID = id
	address.UserID = userID

	err = s.db.Model(&address).Find(&address).Error
	if err == gorm.ErrRecordNotFound {
		l.Warn().Caller().Msg(fmt.Sprintf("client with id '%s' was not found", id))
		return address, e.ErrNotFound
	}
	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	return
}

// Update change specific address by ID.
func (s service) Update(ctx context.Context, id, action string, address entity.Address, clientIDs []string) (err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)
	address.UserID = userID
	address.ID = id

	var clients []entity.Client
	for _, id := range clientIDs {
		clients = append(clients, entity.Client{ID: id})
	}

	fmt.Println("clients")
	fmt.Println(clients)

	switch action {
	case "append":
		fmt.Println("append")
		err = s.db.Debug().Model(&address).Where("id = ? AND user_id = ?", id, userID).Association("Clients").Append(clients)
	case "remove":
		fmt.Println("remove")
		err = s.db.Debug().Model(&address).Where("id = ? AND user_id = ?", id, userID).Association("Clients").Delete(clients)
	default:
		err = s.db.Debug().Model(&address).Where("id = ? AND user_id = ?", id, userID).Updates(address).Error
	}

	if err == gorm.ErrRecordNotFound {
		l.Info().Caller().Msg(fmt.Sprintf("the address with id '%s' and user_id '%s' was not found", address.ID, userID))
		return e.ErrUserNotFound
	}

	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return err
	}
	return nil
}

// Delete disable or delete a specific address by ID.
func (s service) Delete(ctx context.Context, id string) (address entity.Address, err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)
	address.ID = id
	address.UserID = userID

	err = s.db.Debug().Transaction(func(tx *gorm.DB) error {
		if err = tx.Debug().Model(&address).Association("Clients").Clear(); err != nil {
			return err
		}
		if err = tx.Debug().Model(&address).Delete(address).Error; err != nil {
			return err
		}
		return nil
	})

	if err == gorm.ErrRecordNotFound {
		l.Info().Caller().Msg(fmt.Sprintf("the address with id '%s' was not found", id))
		return address, e.ErrUserNotFound
	}

	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	return
}

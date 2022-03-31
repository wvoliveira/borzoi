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
	FindAll(ctx context.Context, search, clientID string, page, limit int) (addresses []entity.Address, err error)
	FindByID(ctx context.Context, id string) (client entity.Client, err error)
	Update(ctx context.Context, payload entity.Address) (address entity.Address, err error)
	Delete(ctx context.Context, id, clientID string) (address entity.Address, err error)

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
		err = query.Where("user_id = ? AND client_id = ?", userID, clientID).Association("Clients").Error
	} else {
		err = s.db.Model(addresses).
			Debug().
			Limit(limit).
			Offset(offset).
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

// FindByID get a specific client by ID.
func (s service) FindByID(ctx context.Context, id string) (client entity.Client, err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)
	client.ID = id

	err = s.db.Model(&client).Where("user_id = ?", userID).Find(&client).Error
	if err == gorm.ErrRecordNotFound {
		l.Warn().Caller().Msg(fmt.Sprintf("client with id '%s' was not found", id))
		return client, e.ErrNotFound
	}
	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	return
}

// Update change specific address by ID.
func (s service) Update(ctx context.Context, address entity.Address) (entity.Address, error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)
	address.UserID = userID

	err := s.db.Debug().Create(&address).Save(&address).Error

	if err == gorm.ErrRecordNotFound {
		l.Info().Caller().Msg(fmt.Sprintf("the address with id '%s' and user_id '%s' was not found", address.ID, userID))
		return address, e.ErrUserNotFound
	}

	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return address, err
	}
	return address, nil
}

// Delete disable or delete a specific address by ID.
func (s service) Delete(ctx context.Context, id, clientID string) (address entity.Address, err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)
	address.ID = id
	address.UserID = userID

	if clientID != "" {
		client := entity.Client{ID: clientID}
		err = s.db.Model(&address).Association("Clients").Delete(client)
	} else {
		err = s.db.Model(&address).Find(&address).Delete(&address).Error
	}

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

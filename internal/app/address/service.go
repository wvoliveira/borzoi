package client

import (
	"context"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/borzoi/internal/pkg/session"
	e "github.com/elga-io/canideos/errors"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
)

// Service encapsulates the link service logic, http handlers and another transport layer.
type Service interface {
	FindAll(ctx context.Context, search string, page, limit int) (clients []entity.Client, err error)
	FindByID(ctx context.Context, id string) (client entity.Client, err error)
	Update(ctx context.Context, id string, payload entity.Client) (link entity.Client, err error)
	Delete(ctx context.Context, id string, delete bool) (link entity.Client, err error)

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
func (s service) FindAll(ctx context.Context, search string, page, limit int) (clients []entity.Client, err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)

	offset := 0
	if page > 1 {
		offset = page * limit
	}

	query := s.db.Model(&entity.Client{}).
		Debug().
		Limit(limit).
		Offset(offset)

	if search != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%[1]s%s%[1]s", "%", search))
	}

	err = query.Where("user_id = ?", userID).Find(&clients).Error

	if len(clients) == 0 {
		l.Warn().Caller().Msg(fmt.Sprintf("clients with search=%s limit=%d offset=%d was not found", search, limit, offset))
		return clients, nil
	}
	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	return
}

// Create add a new client.
func (s service) Create(ctx context.Context, client entity.Client) (err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)
	client.UserID = userID
	client.Active = "true"

	err = s.db.Model(&client).Create(&client).Error
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

// Update change specific client by ID.
func (s service) Update(ctx context.Context, id string, payload entity.Client) (client entity.Client, err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)

	err = s.db.Model(&entity.User{}).Where("user_id = ?", userID).Where("id = ?", id).Updates(&payload).Error
	if err == gorm.ErrRecordNotFound {
		l.Info().Caller().Msg(fmt.Sprintf("the client with id '%s' was not found", id))
		return client, e.ErrUserNotFound
	}

	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	client = payload
	return
}

// Delete disable or delete a specific client by ID.
func (s service) Delete(ctx context.Context, id string, delete bool) (client entity.Client, err error) {
	l := log.Ctx(ctx)
	userID := session.UserGetIDFromContext(ctx)
	client.ID = id
	client.UserID = userID

	if delete {
		err = s.db.Model(&client).Find(&client).Delete(&client).Error
	} else {
		client.Active = "false"
		err = s.db.Model(&client).Updates(&client).Find(&client).Error
	}

	if err == gorm.ErrRecordNotFound {
		l.Info().Caller().Msg(fmt.Sprintf("the client with id '%s' was not found", id))
		return client, e.ErrUserNotFound
	}

	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	return
}

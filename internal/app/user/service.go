package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	e "github.com/elga-io/borzoi/internal/pkg/errors"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Service encapsulates the link service logic, http handlers and another transport layer.
type Service interface {
	FindByID(ctx context.Context, id string) (user entity.User, err error)
	Update(ctx context.Context, id string, payload entity.User) (link entity.User, err error)

	HTTPNew(r *mux.Router)
	HTTPFindAll(w http.ResponseWriter, r *http.Request)
	HTTPFindByID(w http.ResponseWriter, r *http.Request)
	HTTPUpdate(w http.ResponseWriter, r *http.Request)

	HTTPFindMe(w http.ResponseWriter, r *http.Request)
	HTTPUpdateMe(w http.ResponseWriter, r *http.Request)
}

type service struct {
	db    *gorm.DB
	cache *badger.DB
}

// NewService creates a new authentication service.
func NewService(db *gorm.DB, cache *badger.DB) Service {
	return service{db, cache}
}

// FindByID get a shortener link from id.
func (s service) FindByID(ctx context.Context, id string) (user entity.User, err error) {
	l := log.Ctx(ctx)
	user.ID = id

	err = s.db.Model(&user).Preload("Identities").Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		l.Warn().Caller().Msg(fmt.Sprintf("user with id '%s' was not found", id))
		return user, e.ErrNotFound
	}
	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	return
}

// Update change specific user by ID.
func (s service) Update(ctx context.Context, id string, payload entity.User) (user entity.User, err error) {
	l := log.Ctx(ctx)

	err = s.db.Model(&entity.User{}).Where("id = ?", id).Updates(&payload).Error
	if err == gorm.ErrRecordNotFound {
		l.Info().Caller().Msg(fmt.Sprintf("the user with id '%s' was not found", id))
		return user, e.ErrUserNotFound
	}

	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	user = payload
	return
}

// FindMe get user profile from userID.
func (s service) FindMe(ctx context.Context, userID string) (user entity.User, err error) {
	l := log.Ctx(ctx)

	userDB := entity.User{}
	err = s.db.Model(&userDB).Preload("Identities").Where("id = ?", userID).Find(&userDB).Error
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return user, e.ErrAuthUnauthorized
		}
		l.Error().Caller().Msg(err.Error())
		return
	}

	user = userDB
	user.Identities = []entity.Identity{}

	// Set blank password to answer request.
	for _, identity := range userDB.Identities {
		identity.Password = ""
		user.Identities = append(user.Identities, identity)
	}
	return
}

// UpdateMe change self-user by ID.
func (s service) UpdateMe(ctx context.Context, id string, payload entity.User) (err error) {
	l := log.Ctx(ctx)

	err = s.db.Model(&entity.User{}).Where("id = ?", id).Updates(&payload).Error
	if err == gorm.ErrRecordNotFound {
		l.Warn().Caller().Msg(fmt.Sprintf("the user with id '%s' was not found", id))
		return e.ErrUserNotFound
	}

	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return
	}
	return
}

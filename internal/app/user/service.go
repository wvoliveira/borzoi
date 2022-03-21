package user

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	e "github.com/elga-io/canideos/errors"
	"github.com/gorilla/mux"
	zl "github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// Service encapsulates the link service logic, http handlers and another transport layer.
type Service interface {
	FindByID(id string) (user entity.User, err error)
	Update(id string, payload entity.User) (link entity.User, err error)

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
func (s service) FindByID(id string) (user entity.User, err error) {
	user.ID = id

	err = s.db.Debug().Model(&user).Preload("Identities").Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		zl.Warn().Caller().Msg(fmt.Sprintf("user with id '%s' was not found", id))
		return user, e.ErrNotFound
	}
	if err != nil {
		zl.Error().Caller().Msg(err.Error())
		return
	}
	return
}

// Update change specific user by ID.
func (s service) Update(id string, payload entity.User) (user entity.User, err error) {
	t := time.Now()
	payload.UpdatedAt = &t

	err = s.db.Model(&entity.User{}).Where("id = ?", id).Updates(&payload).Error
	if err == gorm.ErrRecordNotFound {
		zl.Info().Caller().Msg(fmt.Sprintf("the user with id '%s' was not found", id))
		return user, e.ErrUserNotFound
	}

	if err != nil {
		zl.Error().Caller().Msg(err.Error())
		return
	}
	user = payload
	return
}

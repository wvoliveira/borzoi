package user

import (
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/gorilla/mux"
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
		msg := fmt.Sprintf("%s warn: the user with id '%s' was not found\n", time.Now(), id)
		return user, errors.New(msg)
	} else if err == nil {
		return
	}
	fmt.Printf("%s error: oh crap, an errors occurred: %s\n", time.Now(), err.Error())
	return
}

// Update change specific user by ID.
func (s service) Update(id string, payload entity.User) (user entity.User, err error) {
	fmt.Printf("%s info: updating user with id '%s'\n", time.Now(), id)

	now := time.Now()
	payload.UpdatedAt = &now

	err = s.db.Model(&entity.User{}).Where("id = ?", id).Updates(&payload).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Printf("%s info: the user with id '%s' was not found", time.Now(), id)
		msg := fmt.Sprintf("%s warn: the user with id '%s' was not found\n", time.Now(), id)
		return user, errors.New(msg)
	} else if err == nil {
		user = payload
		return
	}
	fmt.Printf("%s error: oh crap, an errors occurred: %s", time.Now(), err.Error())
	return
}

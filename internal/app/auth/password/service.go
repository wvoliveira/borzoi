package password

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	e "github.com/elga-io/borzoi/internal/pkg/errors"
	"github.com/elga-io/borzoi/internal/pkg/session"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Service encapsulates the authentication logic.
type Service interface {
	Login(ctx context.Context, identity entity.Identity) (string, error)
	Register(ctx context.Context, identity entity.Identity, user entity.User) error

	HTTPNew(r *mux.Router)
	HTTPLogin(w http.ResponseWriter, r *http.Request)
	HTTPRegister(w http.ResponseWriter, r *http.Request)
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetUID returns the e-mail, google id, facebook id, etc.
	GetUID() string
	// GetRole returns the role.
	GetRole() string
}

type service struct {
	db    *gorm.DB
	cache *badger.DB
}

// NewService creates a new authentication service.
func NewService(db *gorm.DB, cache *badger.DB) Service {
	return service{db, cache}
}

// Login authenticates a user and set a cookie session if authentication succeeds.
func (s service) Login(ctx context.Context, identity entity.Identity) (token string, err error) {
	l := log.Ctx(ctx)

	// Check if identity exists.
	identityDB := entity.Identity{}
	err = s.db.Model(&entity.Identity{}).Where("provider = ? AND uid = ?", identity.Provider, identity.UID).First(&identityDB).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		l.Debug().Caller().Msg(fmt.Sprintf("provider %s + uid %s was not found in database", identity.Provider, identity.UID))
		return token, e.ErrAuthUnauthorized
	}

	if err != nil {
		l.Error().Caller().Msg(fmt.Sprintf("when get identity from database: %s", err.Error()))
		return token, e.ErrInternalServerError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(identityDB.Password), []byte(identity.Password)); err != nil {
		l.Debug().Caller().Msg(e.ErrAuthPasswordDoesntMatch.Error())
		return token, e.ErrAuthUnauthorized
	}

	// Get user info.
	user := entity.User{ID: identityDB.UserID}
	err = s.db.Model(&user).Preload("Identities").Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		l.Warn().Caller().Msg(fmt.Sprintf("user_id %s was not found", identityDB.UserID))
		return token, e.ErrUserNotFound
	}

	if err != nil {
		l.Error().Caller().Msg(err.Error())
		return token, e.ErrInternalServerError
	}

	// Create session in NoSQL.
	token = ksuid.New().String()
	key := fmt.Sprintf(session.DBKey, token)

	err = s.cache.Update(func(txn *badger.Txn) (err error) {
		ee := badger.NewEntry([]byte(key), []byte(user.ID)).WithTTL(time.Hour * 12)
		err = txn.SetEntry(ee)
		return
	})

	if err != nil {
		l.Error().Caller().Msg(err.Error())
	}

	// Update "last login" from user info and I don't care if errors happens.
	now := time.Now()
	identity.LastLogin = &now

	err = s.db.Model(&entity.Identity{}).Where("provider = ? AND uid = ?", identity.Provider, identity.UID).Updates(&identity).Error
	if err != nil {
		l.Error().Caller().Msg(err.Error())
	}

	return token, nil
}

// Register a new user to our database.
func (s service) Register(ctx context.Context, identity entity.Identity, user entity.User) (err error) {
	l := log.Ctx(ctx)

	// Check if identity exists.
	identityDB := entity.Identity{}
	err = s.db.Model(&entity.Identity{}).Where("provider = ? AND uid = ?", identity.Provider, identity.UID).First(&identityDB).Error

	if identityDB.ID != "" {
		l.Info().Caller().Msg(fmt.Sprintf("provider %s and uid %s already exists", identity.Provider, identity.UID))
		return e.ErrAlreadyExists
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Error().Caller().Msg(err.Error())
		return e.ErrInternalServerError
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(identity.Password), 8)
	if err != nil {
		l.Error().Caller().Msg(fmt.Sprintf("when creating a hashed password: %s", err.Error()))
		return fmt.Errorf("internal server error: %s", err.Error())
	}

	identity.ID = ksuid.New().String()
	identity.Password = string(hashedPassword)

	user.ID = ksuid.New().String()
	user.Identities = append(user.Identities, identity)

	err = s.db.Model(&entity.User{}).Create(&user).Error
	return
}

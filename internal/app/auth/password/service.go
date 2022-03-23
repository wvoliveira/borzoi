package password

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/borzoi/internal/pkg/session"
	e "github.com/elga-io/canideos/errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
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
		l.Warn().Caller().Msg(fmt.Sprintf("provider %s + uid %s was not found in database", identity.Provider, identity.UID))
		return token, e.ErrAuthUnauthorized
	}

	if err != nil {
		l.Error().Caller().Msg(fmt.Sprintf("when get identity from database: %s", err.Error()))
		return token, e.ErrInternalServerError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(identityDB.Password), []byte(identity.Password)); err != nil {
		l.Warn().Caller().Msg(e.ErrAuthPasswordDoesntMatch.Error())
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

	token = uuid.New().String()
	key := fmt.Sprintf(session.CacheKey, token)

	err = s.cache.Update(func(txn *badger.Txn) (err error) {
		ee := badger.NewEntry([]byte(key), []byte(user.ID)).WithTTL(time.Hour * 12)
		err = txn.SetEntry(ee)
		return
	})
	return
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
		return errors.New(fmt.Sprintf("internal server error: %s", err.Error()))
	}

	identity.ID = uuid.New().String()
	identity.CreatedAt = time.Now()
	identity.Password = string(hashedPassword)

	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.Identities = append(user.Identities, identity)

	err = s.db.Model(&entity.User{}).Create(&user).Error
	return
}

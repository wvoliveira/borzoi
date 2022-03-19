package password

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// Service encapsulates the authentication logic.
type Service interface {
	Login(ctx context.Context, identity entity.Identity) (string, error)
	Register(ctx context.Context, identity entity.Identity) error

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

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s service) Login(_ context.Context, identity entity.Identity) (token string, err error) {
	fmt.Printf("%s info: initializing login function with identity id %s\n", time.Now(), identity.ID)

	// Check if identity exists.
	identityDB := entity.Identity{}
	err = s.db.Model(&entity.Identity{}).Where("provider = ? AND uid = ?", identity.Provider, identity.UID).First(&identityDB).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("%s warn: this provider + uid was not found in database: %s", time.Now(), err.Error())
		return token, errors.New("sorry, you are not authorized")
	} else if err != nil {
		fmt.Printf("%s error: error when get identity from database: %s", time.Now(), err.Error())
		return token, errors.New("sorry, you are not authorized")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(identityDB.Password), []byte(identity.Password)); err != nil {
		fmt.Printf("%s warn: login failed. Password from payload doesnt match password from database", time.Now())
		return token, errors.New("sorry, you are not authorized")
	}

	// Get user info.
	user := entity.User{}
	err = s.db.Model(&entity.User{}).Where("id = ?", identityDB.UserID).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return token, err
	} else if err != nil {
		return token, err
	}

	token = uuid.New().String()
	key := fmt.Sprintf("users/%s/tokens/%s", user.ID, token)

	err = s.cache.Update(func(txn *badger.Txn) (err error) {
		data, err := json.Marshal(user)
		if err != nil {
			return
		}

		e := badger.NewEntry([]byte(key), []byte(data)).WithTTL(time.Hour * 12)
		err = txn.SetEntry(e)
		return
	})
	return
}

// Register a new user to our database.
func (s service) Register(_ context.Context, identity entity.Identity) (err error) {
	// Check if identity exists.
	identityDB := entity.Identity{}
	err = s.db.Model(&entity.Identity{}).Where("provider = ? AND uid = ?", identity.Provider, identity.UID).First(&identityDB).Error

	if identityDB.ID != "" {
		fmt.Printf("%s debug: provider %s and uid %s already exists\n", time.Now(), identity.Provider, identity.UID)
		return errors.New("this email already exists in our database")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("%s error: unkown error: %s\n", time.Now(), err.Error())
		return errors.New("sorry, internal error happens")
	}

	user := entity.User{}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(identity.Password), 8)
	if err != nil {
		fmt.Printf("%s err: when creating a hashed password: %s\n", time.Now(), err.Error())
		return errors.New(fmt.Sprintf("internal server error: %s", err.Error()))
	}

	identity.ID = uuid.New().String()
	identity.CreatedAt = time.Now()
	identity.Password = string(hashedPassword)

	t := true
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.Role = "user"
	user.Active = &t
	user.Identities = append(user.Identities, identity)

	err = s.db.Model(&entity.User{}).Create(&user).Error
	return
}

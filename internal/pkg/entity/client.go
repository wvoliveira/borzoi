package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Client represents a user info.
type Client struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Score       int    `json:"score"`
	Active      string `json:"active"`

	Phones []Phone `json:"-"`
	Emails []Email `json:"-"`
	Notes  []Note  `json:"-"`

	// Relationship with User and Address model.
	UserID    string     `json:"user_id"`
	Addresses []*Address `json:"addresses" gorm:"many2many:client_addresses;"`
}

func (c *Client) BeforeCreate(_ *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}

// Phone represents phone number with notes.
type Phone struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Note string `json:"note"`
	Num  string `json:"num"`

	// Relationship with Client model.
	ClientID string `json:"client_id"`
}

func (p *Phone) BeforeCreate(_ *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}

// Email represents email with notes.
type Email struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Note    string `json:"note"`
	Address string `json:"address"`

	// Relationship with Client model.
	ClientID string `json:"client_id"`
}

func (e *Email) BeforeCreate(_ *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}

// Note add some text notes for client or any struct.
type Note struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Text string `json:"text"`

	// Relationship with Client model.
	ClientID string `json:"client_id"`
}

func (n *Note) BeforeCreate(_ *gorm.DB) (err error) {
	n.ID = uuid.New().String()
	return
}

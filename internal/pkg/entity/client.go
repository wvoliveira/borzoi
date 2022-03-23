package entity

import (
	"time"
)

// Client represents a user info.
type Client struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name   string `json:"name"`
	Active string `json:"active"`

	Phones []Phone `json:"-"`
	Emails []Email `json:"-"`
	Notes  []Note  `json:"-"`

	// Relationship with User model.
	UserID string `json:"user_id"`
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

type Note struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Text string `json:"text"`

	// Relationship with Client model.
	ClientID string `json:"client_id"`
}

package entity

import (
	"time"
)

// User represents a user info.
type User struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name string `json:"name"`

	Identities []Identity `json:"identities"`
	Clients    []Client   `json:"clients"`
}

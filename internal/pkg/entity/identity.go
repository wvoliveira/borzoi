package entity

import "time"

type Identity struct {
	ID        string     `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	LastLogin *time.Time `json:"last_login"`

	// Phone, email, wechat, github, etc.
	Provider string `json:"provider"`
	// E-mail, username, google id, facebook id, etc.
	UID      string `json:"uid"`
	Password string `json:"password,omitempty"`

	// Relationship with User model.
	UserID string `json:"-"`

	Verified   *bool      `json:"verified" gorm:"default:false"`
	VerifiedAt *time.Time `json:"verified_at"`
}

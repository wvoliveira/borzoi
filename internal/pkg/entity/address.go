package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Address an address for client or any entity.
type Address struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Country    string `json:"country"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	CEP        int    `json:"cep"`
	Street     string `json:"street"`
	Number     int    `json:"number"`
	Complement string `json:"complement"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`

	// Back-reference with User and Client model.
	UserID  string    `json:"-"`
	Clients []*Client `json:"-" gorm:"many2many:client_addresses;"`
}

func (n *Address) BeforeCreate(_ *gorm.DB) (err error) {
	n.ID = uuid.New().String()
	return
}

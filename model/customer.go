package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBOCustomer struct {
	Id             uuid.UUID      `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Name           string         `json:"name" form:"name" gorm:"type:varchar(255)" validate:"required"`
	Address        string         `json:"address" form:"address" gorm:"type:varchar(255)" validate:"required"`
	City           string         `json:"city" form:"city" gorm:"type:varchar(255)" validate:"required"`
	State          string         `json:"state" form:"state" gorm:"type:varchar(255)" validate:"required"`
	PostalCode     int            `json:"postal_code" form:"postal_code" validate:"required"`
	Contact_number int            `json:"contact_number" form:"contact_number" validate:"required"`
	Username       string         `json:"username" form:"username" gorm:"type:varchar(255)" validate:"required"`
	Password       string         `json:"password" form:"password" gorm:"type:varchar(255)" validate:"required"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

func (DBOCustomer) TableName() string { return "customer" }

func (DBOCustomer *DBOCustomer) BeforeCreate(tx *gorm.DB) (err error) {
	DBOCustomer.CreatedAt = time.Now()
	DBOCustomer.UpdatedAt = time.Now()
	return
}

func (DBOCustomer *DBOCustomer) BeforeUpdate(tx *gorm.DB) (err error) {
	DBOCustomer.UpdatedAt = time.Now()
	return
}

type Customer struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	PostalCode     int       `json:"postal_code"`
	Contact_number int       `json:"contact_number"`
	Username       string    `json:"username" `
	Password       string    `json:"password" `
}
type Auth struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
